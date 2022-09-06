package model

import (
	"time"

	"gorm.io/datatypes"
)

// AutoMigrate 初始化数据表
func AutoMigrate(db *Store) error {

	if err := db.AutoMigrate(&User{}, &Book{}, &Tag{}, &Message{}, &Author{}); err != nil {
		return err
	}

	return nil

}

//   username = Column(String(200))
//     password = Column(String(200), default="")
//     salt = Column(String(200))
//     name = Column(String(100))
//     email = Column(String(200))
//     avatar = Column(String(200))
//     admin = Column(Boolean, default=False)
//     active = Column(Boolean, default=True)
//     permission = Column(String(100), default="")
//     create_time = Column(DateTime)
//     update_time = Column(DateTime)
//     access_time = Column(DateTime)

type Book struct {
	ID            int64     `json:"id" gorm:"primaryKey;comment:ID"`
	Title         string    `json:"title" gorm:"type:varchar(255);comment:用户标签列表"`
	Rating        float32   `json:"rating" gorm:"comment:用户标签列表"`
	CTime         time.Time `json:"ctime" form:"ctime" gorm:"column:ctime;autoCreateTime;comment:创建时间"`
	Pubdate       string    `json:"pubdate"`
	Author        string    `json:"author"`
	Authors       []string  `json:"authors"`
	AuthorSort    string    `json:"author_sort"`
	Tags          []string  `json:"tags"`
	Publisher     string    `json:"publisher"`
	Comments      string    `json:"comments"`
	Series        []string  `json:"series"`
	Language      string    `json:"language"`
	ISBN          string    `json:"isbn"`
	Img           string    `json:"img"`
	Thumb         string    `json:"thumb"`
	Collector     string    `json:"collector"`
	CountVisit    int64     `json:"count_visit"`
	CountDownload int64     `json:"count_download"`
	AuthorURL     string    `json:"author_url"`
	PublisherURL  string    `json:"publisher_url"`
	// Files         []File      `json:"files"`
	IsPublic bool `json:"is_public"`
	IsOwner  bool `json:"is_owner"`
}

// {"err": "ok", "kindle_sender": "sender@talebook.org", "book": {"id": 188, "title": "C\u9ad8\u7ea7\u7f16\u7a0b\uff1a\u57fa\u4e8e\u6a21\u5757\u5316\u8bbe\u8ba1\u601d\u60f3\u7684C\u8bed\u8a00\u5f00\u53d1 (C/C++\u6280\u672f\u4e1b\u4e66)", "rating": null, "timestamp": "2022-09-03", "pubdate": "2016-05-01", "author": "\u5409\u661f", "authors": ["\u5409\u661f"], "author_sort": "\u5409\u661f", "tag": "", "tags": [], "publisher": "\u673a\u68b0\u5de5\u4e1a\u51fa\u7248\u793e", "comments": "\u6682\u65e0\u7b80\u4ecb", "series": null, "language": null, "isbn": null, "img": "/get/cover/188.jpg?t=1662179883", "thumb": "/get/thumb_60x80/188.jpg?t=1662179883", "collector": "admin", "count_visit": 0, "count_download": 0, "author_url": "/author/\u5409\u661f", "publisher_url": "/publisher/\u673a\u68b0\u5de5\u4e1a\u51fa\u7248\u793e", "files": [{"format": "EPUB", "size": 678903, "href": "/api/book/188.EPUB"}], "is_public": true, "is_owner": false}, "msg": ""}
// {"err": "ok", "kindle_sender": "sender@talebook.org", "book": {"id": 225, "title": "\u3010\u8c46\u74e3\uff1a8.6\u3011\u300a\u88ab\u8ba8\u538c\u7684\u52c7\u6c14\u300b\uff082015\uff09", "rating": null, "timestamp": "2022-09-03", "pubdate": "2016-04-06", "author": "[\u65e5]\u5cb8\u89c1\u4e00\u90ce, \u53e4\u8d3a\u53f2\u5065", "authors": ["[\u65e5]\u5cb8\u89c1\u4e00\u90ce", "\u53e4\u8d3a\u53f2\u5065"], "author_sort": "[\u65e5]\u5cb8\u89c1\u4e00\u90ce", "tag": "\u751f\u6d3b\u00b7\u52b1\u5fd7", "tags": ["\u751f\u6d3b\u00b7\u52b1\u5fd7"], "publisher": "epub\u638c\u4e0a\u4e66\u82d1", "comments": "\u6682\u65e0\u7b80\u4ecb", "series": null, "language": null, "isbn": "9861371958", "img": "/get/cover/225.jpg?t=1662202092", "thumb": "/get/thumb_60x80/225.jpg?t=1662202092", "collector": "admin", "count_visit": 0, "count_download": 0, "author_url": "/author/[\u65e5]\u5cb8\u89c1\u4e00\u90ce", "publisher_url": "/publisher/epub\u638c\u4e0a\u4e66\u82d1", "files": [{"format": "EPUB", "size": 221154, "href": "/api/book/225.EPUB"}], "is_public": true, "is_owner": false}, "msg": ""}

type Message struct {
	ID     uint           `json:"id" form:"id" gorm:"primaryKey;comment:ID"`
	Title  string         `json:"title" form:"title" gorm:"type:varchar(255); comment:用户标签列表"`
	Status int32          `json:"status" form:"status" gorm:"default:1; comment:用户标签列表"`
	Ctime  time.Time      `json:"ctime" form:"ctime" gorm:"column:ctime;autoCreateTime;comment:创建时间"`
	Data   datatypes.JSON `json:"data" form:"data" gorm:"type:jsonb;comment:创建时间"`
}

type Tag struct {
	ID    int64  `json:"id" form:"id" gorm:"primaryKey;comment:ID"`
	Name  string `json:"name" form:"name" gorm:"type:varchar(100); comment:名称"`
	Count int64  `json:"count" form:"count" gorm:"comment:count"`
}

type Author struct {
	ID      int64  `json:"id" form:"id" gorm:"primaryKey;comment:ID"`
	Name    string `json:"name" form:"name" gorm:"type:varchar(100);comment:名称"`
	Country string `json:"country" form:"country" gorm:"type:varchar(100);comment:名称"`
	Count   int64  `json:"count" form:"count" gorm:"comment:count"`
}


