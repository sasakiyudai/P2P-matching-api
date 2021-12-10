package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	DBMS := "mysql"
	USER := "mysql"
	PASS := "mysql"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "market"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo"

	count := 0
	err := errors.New("")
	DB, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		for {
			if err == nil {
				fmt.Println("")
				break
			}
			fmt.Print(".")
			time.Sleep(time.Second)
			count++
			if count > 180 {
				fmt.Println("")
				fmt.Println("DB connection could not established")
				panic(err)
			}
			DB, err = gorm.Open(DBMS, CONNECT)
		}
	}
	fmt.Println("DB connection established successfully")
}
