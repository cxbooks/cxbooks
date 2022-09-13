package model

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/cxbooks/cxbooks/server/zlog"
	"github.com/cxbooks/epub"
)

type Book struct {
	ID int64 `json:"id" gorm:"primaryKey;comment:ID"`

	Size  int64     `json:"size" gorm:"comment:文件大小"`
	Path  string    `json:"path" gorm:"comment:文件路径"`
	CTime time.Time `json:"ctime" form:"ctime" gorm:"column:ctime;autoCreateTime;comment:创建时间"`
	UTime time.Time `json:"utime" gorm:"comment:文件更新时间"`

	CoverURL string `json:"cover_url" gorm:"comment:封面图片地址"`

	Title string `json:"title" gorm:"type:varchar(255);comment:用户标签列表"`
	// SubTitle represents the EPUB sub-titles.
	SubTitle   string `json:"sub_title,omitempty"`
	Language   string `json:"language"`
	ISBN       string `json:"isbn"`
	Identifier string `json:"identifier"`
	Author     string `json:"author"`
	AuthorURL  string `json:"author_url"`
	AuthorSort string `json:"author_sort"`
	// Publisher identifies the publication's publisher.
	Publisher string `json:"publisher"`
	// Description provides a description of the publication's content.
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags" gorm:"-"` //sqlite 没法存储数组
	// Series is the series to which this book belongs to.
	Series string `json:",omitempty"`
	// SeriesIndex is the position in the series to which the book belongs to.
	SeriesIndex string `json:",omitempty"`
	PublishDate string `json:"pubdate"`

	Rating float32 `json:"rating" gorm:"comment:用户标签列表"`

	PublisherURL string `json:"publisher_url"`

	CountVisit    int64 `json:"count_visit"`
	CountDownload int64 `json:"count_download"`
}

// Save 存储图书元数据到数据库
func (book *Book) Save(store *Store) error {
	//TODO before save data
	return store.Save(book).Error

}

// GetMetadataFromFile reads metadata from an epub file.
func (book *Book) GetMetadataFromFile(kv *KV) error {

	_, err := os.Stat(book.Path)
	if os.IsNotExist(err) {
		// path/to/whatever does not exist
		zlog.D(`文件不存在`, book.Path, `无法访问或者不存在`)
		return errors.New(`缓存目录无法访问或者不存在`)
	}

	e, err := epub.Open(book.Path)
	if err != nil {
		zlog.E(`打开文件失败：`, err.Error())
		return err
	}
	defer e.Close()

	opf, err := e.Package()
	if err != nil {
		zlog.E(`解析文件,`, book.Path, `失败：`, err.Error())
		return err
	}

	book.parseOPF(opf)

	//解析出来封面信息
	if book.CoverURL != `` {

		fp, err := e.OpenItem(book.CoverURL)

		if err != nil {
			zlog.E(`解析文件,`, book.Path, `失败：`, err.Error())
			return err
		}

		coverURL := `uuid` + book.CoverURL //TODO 这里要用bookid 获取其他标记避免冲突

		err = kv.Write(coverURL, fp, 0)
		if err != nil {
			zlog.E(`存储封面失败,`, book.Path, `失败：`, err.Error())
			return err
		}
		//将 CoverURL 地址覆盖成解析后的地址
		book.CoverURL = coverURL

	}

	return nil
}

func (m *Book) parseOPF(opf *epub.PackageDocument) {

	mdata := opf.Metadata

	m.Language = elt2FirstStr(mdata.Language)
	m.Tags = elt2str(mdata.Subject)
	m.Description = elt2FirstStr(mdata.Description)
	m.Publisher = elt2FirstStr(mdata.Publisher)

	//TODO get uuid
	// for _, id := range mdata.Identifier {
	// 	m.Identifier = append(m.Identifier, Identifier{
	// 		Value:  id.Value,
	// 		Scheme: id.Scheme,
	// 	})
	// }

	if len(mdata.Creator) > 0 {
		m.Title = mdata.Title[0].Value
	} else {
		zlog.W(`查找图书名失败，使用文件名作为标题`, m.Path)
		fileName := filepath.Base(m.Path)
		m.Title = fileName[:len(fileName)-len(filepath.Ext(fileName))]
	}

	if len(mdata.Creator) > 0 {
		m.Author = mdata.Creator[0].Value
	}

	if len(mdata.Date) > 0 {
		m.PublishDate = mdata.Date[0].Value
	}

	m.parseMeta(opf)

}

func (m *Book) parseMeta(opf *epub.PackageDocument) {

	metas := opf.Metadata.Meta
	for _, meta := range metas {
		switch meta.Name {
		case "calibre:series":
			m.Series = meta.Content

		case "calibre:series_index":
			m.SeriesIndex = meta.Content
		case "cover":
			id := meta.Content

			if opf.Manifest != nil {
				items := opf.Manifest.Items
				for i := len(items) - 1; i >= 0; i-- {
					if items[i].ID == `cover-image` {
						m.CoverURL = items[i].Href
					}
				}

			}
			println(id)
		}

	}

}

func elt2str(elt []epub.Element) []string {
	s := make([]string, len(elt))

	for i, e := range elt {
		s[i] = e.Value
	}

	return s
}

func elt2FirstStr(elt []epub.Element) string {

	if len(elt) > 0 {
		return elt[0].Value
	}
	return ""

}
