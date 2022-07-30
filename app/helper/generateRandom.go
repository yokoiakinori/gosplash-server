package helper

import (
	"crypto/rand"
	"errors"
)

func MakeRandomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("乱数の生成に失敗しました。")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}

	return result, nil
}