package model

type AccessToken struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Token string `xorm:"unique varchar(255)" form:"token" json:"token"`
	RefreshToken string `xorm:"unique varchar(255)" form:"refresh_token" json:"refresh_token"`
	ExpiresAt string `xorm:"varchar(255)" form:"expires_at" json:"expires_at"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
}