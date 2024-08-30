package models

type Auth struct {
	ID       int32  `json:"id" gorm:"type:uuid;primary_key;autoIncrement"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}
