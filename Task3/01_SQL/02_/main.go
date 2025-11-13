package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type user_mark struct{
	gorm.Model

	Name string
	Email *string
	Age uint8
	CreatedAt time.Time
	UpdateAt time.Time

}

type connectInfo struct{
	username string
	password string
	host string
	port string
	dbName string
}



func main() {
	fmt.Println("hello world")
}