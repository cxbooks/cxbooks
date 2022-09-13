package model

import (
	"time"

	"github.com/cxbooks/cxbooks/server/zlog"
	"gorm.io/gorm/clause"
)

// AutoMigrate 初始化数据表
func AutoMigrate(db *Store) error {

	if err := db.AutoMigrate(&User{}, &Book{}, &Tag{}, &Session{}, &Message{}, &Author{}); err != nil {
		return err
	}

	//createDefaultData
	zlog.I(`创建初始化管理员账号，如果已经有改账号则忽略`)
	return initDefaultData(db)

}

func initDefaultData(db *Store) error {

	//创建默认用户
	admin := &User{
		UserID:     1,
		Email:      "",
		Mobile:     "",
		UserName:   "admin",
		NickName:   "admin",
		AvatarUrl:  "",
		Language:   "zh-CN",
		Tags:       "",
		WxUnionID:  "",
		RoleID:     1,
		Credential: encryptPassword(`admin123`),
		Locked:     false,
		MFASwitch:  false,
		CTime:      time.Now(),
		UTime:      time.Now(),
		ATime:      time.Now(),
	}

	return db.Clauses(clause.OnConflict{
		// Columns:   []clause.Column{{Name: "key"}},
		DoNothing: true,
	}).Create(admin).Error

}
