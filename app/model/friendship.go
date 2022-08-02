package model

type Friendship struct {
	Id int64 `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	ApplicantId int64 `xorm:"not null" form:"applicant_id" json:"applicant_id"`
	AuthorizerId int64 `xorm:"not null" form:"authorizer_id" json:"authorizer_id"`
	Created string `xorm:"timestamp created" form:"created" json:"created"`
	Updated string `xorm:"timestamp updated" form:"updated" json:"updated"`
}