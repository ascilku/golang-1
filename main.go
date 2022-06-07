package main

import (
	"api18/handler"
	"api18/member"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_inready?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		var valueMember []member.Member
		db.Find(&valueMember)

		newRepository := member.NewRepository(db)
		newServe := member.NewServeMember(newRepository)
		newHandlerMember := handler.NewHandlerMember(newServe)

		routes := gin.Default()
		routes.POST("member", newHandlerMember.SaveHandlerMember)
		routes.Run()
	}
}
