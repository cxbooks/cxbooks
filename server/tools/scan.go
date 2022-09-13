// scan 目录扫描搜刮工具
package tools

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/cxbooks/cxbooks/server/model"
	"github.com/cxbooks/cxbooks/server/zlog"
)

// 文件扫描器
type Scanner struct {
	kv    *model.KV
	store *model.Store
}

func OpenScanner(path string) (*Scanner, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.D(`缓存目录`, path, `无法访问或者不存在 `, err)
		return nil, errors.New(`缓存目录无法访问或者不存在`)
	}

	scanner := &Scanner{}

	scanner.kv, err = model.OpenKV(filepath.Join(path, `scanner.db`))

	return scanner, err
}

func (m *Scanner) Close() {
	if m.kv != nil {
		zlog.I(`退出扫描缓存DB`)
		m.kv.Close()
	}
}

func (m *Scanner) Save(book *model.Book) error {

	err := book.GetMetadataFromFile(m.kv)

	if err != nil {
		zlog.E(`获取图书基本元素失败 `, err.Error())
		return err
	}

	//存储图书到数据库

	if err := book.Save(m.store); err != nil {
		zlog.E(`存储图书失败：`, err)
		return err
	}

	return nil

}

// Scan 扫描路径下epub文件
func Scan(root string, depth int) <-chan *model.Book {
	file := make(chan *model.Book)
	zlog.I(`开始扫描路径:`, root)
	go func() {

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				zlog.E(`扫描路径失败：`, path)
				return err
			}

			if !info.IsDir() && filepath.Ext(info.Name()) == `.epub` {

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
