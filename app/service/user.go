package service

import (
	"gosplash-server/app/model"
)

type UserService struct {
	
}

func (UserService) Register(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (UserService) Login(email string, password string) error {
	user := model.User{}
	_, err := DbEngine.Where("email", email).Get(&user)
	if user.Password !== password {
		err = "パスワードが正しくありません。"
		return err
	}
	
	if err != nil {
		return err
	}
	return nil
}