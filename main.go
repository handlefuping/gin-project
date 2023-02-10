package main

import (
	"gin-project/controller"
	"gin-project/middleware"
	"gin-project/model"
	"gin-project/util"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)


func init() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db open fail")
	}
	model.DBInstance = db
	db.AutoMigrate(&model.Video{})
	db.AutoMigrate(&model.User{})
}

func main() {
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 媒体静态文件服务
	r.Static(util.GetVideoStoreDir(), util.GetVideoStoreDir())

	userController := controller.User{}
	user := r.Group("/user")

	user.POST("/login", userController.LoginUser)
	user.POST("/register", userController.RegisterUser)
	user.Use(middleware.AuthRequired())

	{
		user.GET("/:id", userController.GetUserById)
		user.POST("/upload", userController.Upload)
	}

	r.Run()
}