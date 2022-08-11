package model

type Comment struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	UserId int64 `xorm:"not null" form:"user_id" json:"user_id"`
	PostId int64 `xorm:"not null" form:"post_id" json:"post_id"`
	Content string `xorm:"varchar(255)" form:"content" json:"content"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
	Version int `xorm:"'version'" form:"updated" json:"updated"`
}