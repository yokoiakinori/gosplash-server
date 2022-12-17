package model

type User struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Name string `xorm:"varchar(60)" form:"name" json:"name"`
	Email string `xorm:"unique varchar(255)" form:"email" json:"email"`
	Password string `xorm:"varchar(255)" form:"password" json:"password"`
	Description string `xorm:"varchar(255)" form:"description" json:"description"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
}

type Register struct {
	Name string `xorm:"varchar(60)" json:"name" json:"name"`
	Email string `xorm:"unique varchar(255)" json:"email" json:"email"`
	Password string `xorm:"varchar(255)" json:"password" json:"password"`
}

type UserNameAndIcon struct {
	Name string `xorm:"varchar(60)" form:"name" json:"name"`
	Icon IconPath `xorm:"extends" from:"icon" json:"icon"`
}