package setup

import (
	"fmt"
	
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
}