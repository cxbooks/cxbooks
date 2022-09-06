package model

import (
	"time"

	"github.com/cxbooks/cxbooks/zlog"
)

const cacheKeySession = `user:sessions:`

type Session struct {
	Session      string        `json:"session" gorm:"type:char(40);primaryKey;comment:当前会话id"`
	UserID       string        `json:"user_id" gorm:"type:char(36);default:null;comment:用户ID(UUID)"`
	RemoteIP     string        `json:"remote_ip" gorm:"column:remote_ip;type:cidr;comment:用户登录ip"`
	UserAgent    string        `json:"user_agent" gorm:"column:user_agent;default:null;type:varchar(255);comment:客户端平台标识"`
	ClientDevice string        `json:"client_device" gorm:"column:client_device;type:varchar(50);comment:客户端登录设备"`
	DistrictId   int32         `json:"district_id" gorm:"column:district_id;type:int8;comment:登录城市id"`
	FromUrl      string        `json:"from_url" gorm:"column:from_url;type:varchar(255);comment:登录上一跳服务"`
	Duration     time.Duration `json:"duration" gorm:"type:int8;comment:会话有效时间"`
	CTime        time.Time     `json:"ctime" form:"ctime" gorm:"column:ctime;autoCreateTime;comment:创建时间"`
	UTime        time.Time     `json:"utime" form:"utime" gorm:"column:utime;autoUpdateTime;comment:更新时间"`
	UserInfo     User          `json:"user_info" gorm:"-"`
}

func (m *Session) Clean(store *Store) {
	//这里忽略错误，由系统其他goroutine 定时清理过期session
	if m.Session != `` {
		store.Delete(m)
		key := cacheKeySession + m.Session
		store.LRU.Remove(key)
	}

}

// parseSessionUser 获取session 关联的用户信息，
func FirstSessionByID(store *Store, id string) (*Session, error) {

	key := cacheKeySession + id

	if raw, err := store.LRU.Get(key); err == nil {
		zlog.D(`hit cachekey: `, key)
		if ret, ok := raw.(*Session); ok {
			return ret, nil
		}

	}
	//get session from db
	ret := &Session{}

	if err := store.Table(`sessions`).Where(`id = ?`, id).First(&ret).Error; err != nil {
		zlog.E(`查询session：`, id)
		return ret, err
	}

	//get session related user
	err := store.Table(`users`).Where(`user_id = ?`, ret.UserID).First(&ret.UserInfo).Error
	if err != nil {
		zlog.E(`查询session 关联的UserID失败`, ret.UserID)
	}

	return ret, err

}
