// scan 目录扫描搜刮工具
package tools

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
)

// const MaxThread = 1

type ScanStatus struct {
	Running bool     `json:"running"`
	Count   int      `json:"count"`
	Errs    []string `json:"errs"`
}

type ScannerManager struct {
	kv    *model.KV
	store *model.Store

	ctx context.Context

	sync.RWMutex

	scan *Scanner
}

// 文件扫描器
type Scanner struct {
	Count int      `json:"count"` //扫描文件计数
	Errs  []string `json:"errors"`

	//扫描错误详情信息
	root string
	wg   *sync.WaitGroup
}

func NewScannerManager(ctx context.Context, dbpath string, store *model.Store) (*ScannerManager, error) {

	scanner := &ScannerManager{
		store: store,
		ctx:   ctx,
	}

	zlog.I(`初始化扫描管理`)

	_, err := os.Stat(dbpath)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.D(`缓存目录`, dbpath, `无法访问或者不存在 `, err)
		return nil, errors.New(`缓存目录无法访问或者不存在`)
	}

	scanner.kv, err = model.OpenKV(filepath.Join(dbpath, `scanner.db`))

	return scanner, err
}

func (m *ScannerManager) IsRunning() bool {
	m.RLock()
	defer m.RUnlock()
	return m.scan != nil
}

func (m *ScannerManager) Status() ScanStatus {
	s := ScanStatus{
		Running: true,
		Count:   0,
		Errs:    []string{},
	}

	s.Running = m.IsRunning()
	if s.Running {
		s.Count = m.scan.Count
		s.Errs = m.scan.Errs
	}

	return s

}

func (m *ScannerManager) Start(path string, maxThread int) {

	if m.IsRunning() {
		zlog.E(`扫描器已经在执行，放弃执行`)
		return
	}
	zlog.I(`启动文件目录扫描`)
	m.NewScan(path, maxThread)

}

func (m *ScannerManager) NewScan(path string, maxThread int) {
	m.Lock()
	m.scan = NewScan(path, maxThread) //new scanner
	m.Unlock()

	m.scan.Start(m.ctx, m.kv, m.store, maxThread)
}

func (m *ScannerManager) Stop() {

	m.Lock()
	defer m.Unlock()

	if m.scan != nil {
		m.scan.Stop()
	}

}

func (m *ScannerManager) Close() {
	if m.kv != nil {
		zlog.I(`退出扫描缓存DB`)
		m.kv.Close()
	}
}

func (m *Scanner) Stop() {
	if m.wg != nil {
		m.wg.Wait()
	}
}

func NewScan(root string, concurrent int) *Scanner {

	return &Scanner{
		Count: 0,
		root:  root,
		wg:    new(sync.WaitGroup),
		Errs:  []string{},
	}

}

func (m *Scanner) Start(ctx context.Context, kv *model.KV, store *model.Store, maxThread int) {

	books := m.Walk(ctx)

	concurrent := make(chan struct{}, maxThread)

	errChan := make(chan string)

	go func() {

		defer func() {
			close(errChan)
			zlog.I(`退出扫描程序`)
		}()

		for {

			select {

			case <-ctx.Done():
				zlog.W(`接收到退出命令，退出扫描`)
				return
			case book, ok := <-books:

				if !ok {
					//books is closed
					zlog.I(`书籍为空，退出解析文件。`)

					return
				}

				m.wg.Add(1)
				zlog.D(`开始解析: `, book.Path, ` 到数据库`)
				concurrent <- struct{}{}
				go func(b *model.Book) {

					defer m.wg.Done()

					if err := b.GetMetadataFromFile(); err != nil {
						zlog.E(`获取图书基本元素失败 `, err.Error())
						errChan <- `文件：` + b.Path + ` 解析失败：` + err.Error()

					} else {

						//存储封面图片数据
						if err := kv.Set(b.CoverURL, b.GetCoverData(), 0); err != nil {
							zlog.E(`存储封面失败,`, book.Path, `失败：`, err.Error())
							// return err
						}

						//存储图书到数据库 TODO 处理重复问题
						if err := b.Save(store); err != nil {
							zlog.E(`存储图书失败：`, err)
							errChan <- `文件：` + b.Path + ` 存储失败：` + err.Error()
						}
					}

					<-concurrent

				}(book)

			}
		}

	}()

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		for err := range errChan {
			m.Errs = append(m.Errs, err)
		}

	}()

}

// Scan 扫描路径下epub文件
func (m *Scanner) Walk(ctx context.Context) <-chan *model.Book {
	file := make(chan *model.Book)
	zlog.I(`开始扫描路径:`, m.root)
	go func() {

		err := filepath.Walk(m.root, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				zlog.E(`扫描路径失败：`, path)
				return err
			}

			if !info.IsDir() && filepath.Ext(info.Name()) == `.epub` {
				zlog.I(`扫描的到文件：`, path)
				book := &model.Book{
					ID:            0,
					Size:          info.Size(),
					Path:          path,
					CTime:         time.Now(),
					UTime:         time.Now(),
					Rating:        0,
					PublishDate:   `1970-01-01`,
					CountVisit:    0,
					CountDownload: 0,
				}

				file <- book
				m.Count++
			}
			return nil
		})

		if err != nil {
			//TODO 如果扫描失败应该让同步返回到操作
			zlog.E(`扫描路径失败：`, err.Error())
		}
		close(file)
	}()

	return file
}
