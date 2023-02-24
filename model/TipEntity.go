package model

type TipEntity map[string]interface{}

func (t *TipEntity) TableName() string {
	return "tip"
}
