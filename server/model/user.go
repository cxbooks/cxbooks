package model

import "time"

type User struct {
	UserID     string    `json:"user_id" form:"user_id" gorm:"type:char(36);primaryKey;comment:用户ID(UUID)"`
	Email      string    `json:"email" form:"email"  gorm:"type:varchar(100);comment:用户邮箱"`
	Mobile     string    `json:"mobile" form:"mobile"   gorm:"type:varchar(50);comment:用户手机"`
	UserName   string    `json:"user_name" form:"user_name" gorm:"uniqueIndex:idx_users_user_name;type:varchar(100);comment:用户名"`
	NickName   string    `json:"nick_name" form:"nick_name" gorm:"type:varchar(100); comment:用户昵称"`
	AvatarUrl  string    `json:"avatar_url" form:"avatar_url" gorm:"type:varchar(255); comment:用户头像图片地址"`
	Language   string    `json:"language" form:"language" gorm:"type:varchar(20); comment:用户默认语言选项"`
	Tags       string    `json:"tags" form:"tags" gorm:"type:varchar(255); comment:用户标签列表"`
	WxUnionID  string    `json:"wx_unionid" form:"wx_unionid" gorm:"type:varchar(50); comment:微信unionID"`
	RoleID     int64     `json:"role_list" gorm:"-"`
	Credential string    `json:"credential" form:"credential" binding:"required" gorm:"not null;type:varchar(100); comment:加密密码"`
	Locked     int32     `json:"locked" form:"locked" gorm:"default:2;comment:用户是否被锁定"`
	MFASwitch  bool      `json:"mfa_switch" form:"mfa_switch" gorm:"comment:mfa虚拟认证"`
	CTime      time.Time `json:"ctime" form:"ctime" gorm:"column:ctime;autoCreateTime;comment:创建时间"`
	UTime      time.Time `json:"utime" form:"utime" gorm:"column:utime;autoUpdateTime;comment:更新时间"`
	ATime      time.Time `json:"atime" form:"uatime" gorm:"column:atime;autoCreateTime;comment:访问时间"`
}
