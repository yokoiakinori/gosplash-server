package service

import (
	"errors"

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
	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if user.Password != password {
		err = errors.New("パスワードが一致しません。")
	}
	
	if err != nil {
		return err
	}
	return nil
}