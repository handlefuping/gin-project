package main

import (
	"fmt"
	"gin-project/controller"
	"gin-project/middleware"
	"gin-project/model"
	"gin-project/util"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func initViper()  {
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read config error")
	}
}

func initDB() {
	dsn :=  fmt.Sprintf("%s:%s@tcp(%s:%d)/gin?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port") )
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db open fail")
	}
	model.DBInstance = db
	db.AutoMigrate(&model.Media{})
	db.AutoMigrate(&model.User{})
}

func initMedia()  {
	err := os.Mkdir(viper.GetString("media.dir"), 0750)
	log.Println(err.Error(), "视频文件夹初始化")
}

func init() {
	initViper()

	initDB()
	initMedia()
}

func main() {
	r := gin.New()
	r.MaxMultipartMemory = viper.GetInt64("video.limitSize")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 媒体静态文件服务
	r.Static(util.GetMediaStoreDir(), util.GetMediaStoreDir())

	userController := controller.User{}
	user := r.Group("/user")

	user.POST("/login", userController.LoginUser)
	user.POST("/register", userController.RegisterUser)
	user.Use(middleware.AuthRequired())

	{
		user.GET("/:id", userController.GetUserById)
		user.POST("/upload", userController.UploadMedia)
		user.POST("/public", userController.PublicMedia)
		user.POST("/ban", userController.BanMedia)
	}

	r.Run()
}