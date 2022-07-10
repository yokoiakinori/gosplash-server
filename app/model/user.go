package model

type User struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	name string `xorm:"varchar(60)" form:"name" json:"name"`
	email string `xorm:"unique varchar(255)" form:"email" json:"email"`
	password string `xorm:"varchar(255)" form:"password" json:"password"`
	description string `xorm:"varchar(255)" form:"description" json:"description"`
	created string `xorm:"timestamp created" form:"created" json:"created"`
	updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
}