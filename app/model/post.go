package model

type Post struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Name string `xorm:"not null" form:"name" json:"name"`
	Path string `xorm:"varchar(60)" form:"path" json:"path"`
	UserId int64 `xorm:"not null" form:"user_id" json:"user_id"`
	Description string `xorm:"varchar(255)" form:"description" json:"description"`
	ViewCount int64 `xorm:"default 0" form:"view_count" json:"view_count"`
	DownloadCount int64 `xorm:"default 0" form:"download_count" json:"download_count"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
}

type PostListAndUser struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Name string `xorm:"not null" form:"name" json:"name"`
	Path string `xorm:"varchar(60)" form:"path" json:"path"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
	User UserNameAndIcon `xorm:"extends" from:"user" json:"user"`
}