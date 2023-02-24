package model

type TopicEntity struct {
	BaseEntity
	Content string `gorm:"content" json:"content"` // 正文内容
	Title   string `gorm:"title" json:"title"`     // 正文标题
	Type    string `gorm:"type" json:"type"`       // 帖子类型
	Role    string `gorm:"role" json:"userRole"`   // 用户身份
	Time    string `gorm:"time" json:"time"`
}

func (t *TopicEntity) TableName() string {
	return "topic"
}

type CommentEntity struct {
	Id      int64  `json:"id" gorm:"id;primaryKey;autoIncrement;comment:主键id"`
	Time    string `gorm:"time" json:"time"`
	Content string `gorm:"content" json:"content"` // 评论文字内容
	Target  string `gorm:"target" json:"target"`   // 目标帖子id
}

func (t *CommentEntity) TableName() string {
	return "comment"
}

type TopicGetEntity struct {
	Id   int64  `json:"id" gorm:"id;primaryKey;autoIncrement;comment:主键id"`
	Time string `gorm:"time" json:"time"`
	//UpdatedAt LocalTime      `json:"updatedAt" gorm:"updated_at;comment:更新时间"`
	//DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at;comment:删除时间"` // 查询这个字段但是不返回这个字段
	Content string          `gorm:"content" json:"content"` // 正文内容
	Title   string          `gorm:"title" json:"title"`     // 正文标题
	Type    string          `gorm:"type" json:"type"`       // 帖子类型
	Role    string          `gorm:"role" json:"userRole"`   // 用户身份
	Comment []CommentEntity `gorm:"-" json:"commentList"`
}

func (t *TopicGetEntity) TableName() string {
	return "topic"
}
