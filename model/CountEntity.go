package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Tag []string

// 项目分数节点
type CountModelEntity struct {
	Id     int64   `json:"id" gorm:"id;primaryKey;autoIncrement;comment:主键id"`
	Label  string  `gorm:"label" json:"label"`   // 分数的中文名
	Top    float64 `gorm:"top" json:"top"`       // 上限，为`-1`表示无上限
	Unique bool    `gorm:"unique" json:"unique"` // 唯一性
	Index  string  `gorm:"index" json:"index"`
}

type CountEntity struct {
	Applytime    string  `gorm:"applytime" json:"applyTime"`       // 申请时间戳，若不可申请，此项为`null`
	Deadlinetime string  `gorm:"deadlinetime" json:"deadLineTime"` // 申请时间戳，若不可申请，此项为`null`
	Content      Tag     `gorm:"content" json:"content"`           // 具体内容数组，累计分数的项目列表，若`unique`为真，则数组长度为`0`或`1`
	Index        string  `gorm:"index" json:"index"`               // 分数的索引
	Label        string  `gorm:"label" json:"label"`               // 分数的中文名
	Top          float64 `gorm:"top" json:"top"`                   // 上限，为`-1`表示无上限
	Unique       bool    `gorm:"unique" json:"unique"`             // 唯一性
	Value        float64 `gorm:"value" json:"value"`               // 分数
	Countid      string  `gorm:"countid" json:"countid"`
	Teamyear     int     `gorm:"teamyear" json:"year"`
}

func (t *CountEntity) TableName() string {
	return "count"
}

func (t *CountModelEntity) TableName() string {
	return "countmodel"
}

type SonCountEntity struct {
	Name    string  `gorm:"name" json:"name"`
	Value   float64 `gorm:"value" json:"value"`
	Content Tag     `gorm:"content" json:"content"`
}

func (t *SonCountEntity) TableName() string {
	return "sunCount"
}

func (t *Tag) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

func (t Tag) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type Reasons struct {
	Content string `gorm:"content" json:"reason"`
}

func (t *Reasons) TableName() string {
	return "reasonpub"
}
