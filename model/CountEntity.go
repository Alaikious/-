package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Tag []string

// 项目分数节点
type CountModelEntity struct {
	Id        int64              `json:"id" gorm:"id;primaryKey;autoIncrement;comment:主键id"`
	ApplyTime string             `gorm:"applyTime" json:"applyTime"` // 申请时间戳，若不可申请，此项为`null`
	Content   Tag                `gorm:"content" json:"content"`     // 具体内容数组，累计分数的项目列表，若`unique`为真，则数组长度为`0`或`1`
	Index     string             `gorm:"index" json:"index"`         // 分数的索引
	Label     string             `gorm:"label" json:"label"`         // 分数的中文名
	List      []CountModelEntity `gorm:"-" json:"list"`              // 子分数，若为节点，则不为`null`
	Top       float64            `gorm:"top" json:"top"`             // 上限，为`-1`表示无上限
	Unique    bool               `gorm:"unique" json:"unique"`       // 唯一性
	Value     float64            `gorm:"value" json:"value"`         // 分数
	Realindex int                `gorm:"realindex" json:"realindex"`
}

type CountEntity struct {
	ApplyTime string        `gorm:"applyTime" json:"applyTime"` // 申请时间戳，若不可申请，此项为`null`
	Content   Tag           `gorm:"content" json:"content"`     // 具体内容数组，累计分数的项目列表，若`unique`为真，则数组长度为`0`或`1`
	Index     string        `gorm:"index" json:"index"`         // 分数的索引
	Label     string        `gorm:"label" json:"label"`         // 分数的中文名
	List      []CountEntity `gorm:"-" json:"list"`              // 子分数，若为节点，则不为`null`
	Top       float64       `gorm:"top" json:"top"`             // 上限，为`-1`表示无上限
	Unique    bool          `gorm:"unique" json:"unique"`       // 唯一性
	Value     float64       `gorm:"value" json:"value"`         // 分数
	Countid   string        `gorm:"countid" json:"countid"`
	Realindex int           `gorm:"realindex" json:"realindex"`
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
	Reason Tag `gorm:"reason" json:"reason"`
}

func (t *Reasons) TableName() string {
	return "reason"
}
