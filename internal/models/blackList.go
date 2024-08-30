package models

type BlackList struct {
	ID int32 `json:"id" gorm:"type:uuid;primary_key;autoIncrement"`
	Ip string
}
