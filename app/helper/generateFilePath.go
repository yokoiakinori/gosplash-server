package helper

import (
	"time"
)

func MakeFilePath(directory string) (string, error) {
	var randomDigit uint32 = 25
	randomStr, _ := MakeRandomStr(randomDigit)
	now := time.Now()
	nowStr := now.Format("2006-01-0-03-04-05") // Goのフォーマットは独特 決まった日付の例じゃないと動かない

	return directory + "/" + nowStr + "-" + randomStr, nil
}