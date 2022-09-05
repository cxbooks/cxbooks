package model

import "gorm.io/gorm"

// AutoMigrate 初始化数据表
func AutoMigrate(db *gorm.DB) error {

	if err := db.AutoMigrate(&User{}, &Book{}, &Tag{}, &Author{}); err != nil {
		return err
	}

	return nil

}

type User struct {
	UserName string
	NickName string
	Email    string
}

type Book struct {
	ID int64
}

type Tag struct {
	BookID int64
	Name   string
}

type Author struct {
	Name string
}
