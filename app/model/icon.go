package model

type Icon struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Path string `xorm:"varchar(60)" form:"path" json:"path"`
	UserId int64 `xorm:"not null" form:"user_id" json:"user_id"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
}

type IconPath struct {
	Path string `xorm:"varchar(60)" form:"path" json:"path"`
}