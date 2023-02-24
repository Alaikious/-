package model

// RecordEntity 申报记录
type RecordEntity struct {
	Id       int64  `json:"ids" gorm:"id;primaryKey;autoIncrement;comment:主键id"`
	Complain string `gorm:"complain" json:"complain"`   // 申诉内容
	Content  string `gorm:"content" json:"content"`     // 文字描述
	Files    Tag    `gorm:"files" json:"files"`         // 图片地址数组
	CountID  string `gorm:"countid" json:"id"`          // 申报id
	Index    string `gorm:"index" json:"index"`         // 申报项目代号，通过层级组合生成; `v-2-2`表示**德育素质-纪实加减分-社会责任记实分**; `v-1`表示**德育素质-基本评定分**
	Label    string `gorm:"label" json:"label"`         // 中文名
	Reason   string `gorm:"Reason" json:"rejectReason"` // 驳回理由
	State    string `gorm:"state" json:"state"`         // 状态
	Time     string `gorm:"time" json:"time"`           // 申报时间
	Value    int64  `gorm:"value" json:"value"`         // 分数
}

func (t *RecordEntity) TableName() string {
	return "record"
}

type ApplyEntity struct {
	Name      string `gorm:"name" json:"name"`
	StartTime string `gorm:"starttime" json:"startTime"`
	EndTime   string `gorm:"endtime" json:"endTime"`
	State     string `gorm:"state" json:"state"`
}

func (t *ApplyEntity) TableName() string {
	return "apply"
}

type SuggestEntity struct {
	Content string `gorm:"content" json:"content"` // 建议内容
	Countid string `gorm:"countid" json:"countid"` // 是否匿名
	Target  string `gorm:"target" json:"target"`   // 辅导员工号
}

func (t *SuggestEntity) TableName() string {
	return "suggest"
}
