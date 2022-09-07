package model

import "time"

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
