package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/datatypes"
)

type Message struct {
	ID     int64          `json:"id" form:"id" gorm:"primaryKey;comment:ID"`
	UserID int64          `json:"user_id" form:"user_id" gorm:"comment:用户ID(UUID)"`
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

func FindMessages(store *Store, userID int64, limit int) ([]Message, error) {

	messages := []Message{}

	err := store.Table(`messages`).Where(`user_id = ?`, userID).Limit(limit).Find(&messages).Error

	return messages, err

}

// JSON defiend JSON data type, need to implements driver.Valuer, sql.Scanner interface
type JSON json.RawMessage

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// MarshalJSON to output non base64 encoded []byte
func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

// UnmarshalJSON to deserialize []byte
func (j *JSON) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*j = JSON(result)
	return err
}

// GormDataType gorm common data type
func (JSON) GormDataType() string {
	return "json"
}
