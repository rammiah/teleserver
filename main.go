package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"os"
)

var (
	db *gorm.DB
)

func main() {
	// 读取配置文件
	file, err := os.Open("config.json")
	if err != nil {
		panic("no config file or open failed")
	}
	data, _ := ioutil.ReadAll(file)
	file.Close()
	var config = struct {
		DBUrl string `json:"dbUrl"`
	}{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic("no db url config")
	}
	// 连接数据库
	db, err = gorm.Open("mysql", config.DBUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 关闭gorm的log
	//db.LogMode(false)

	gin.DisableConsoleColor()
	engine := gin.Default()
	// 设置路由信息
	engine.POST("/userRegister", userRegister)
	engine.GET("/userLogin", userLogin)
	engine.GET("/ping", ping)
	engine.GET("/getMenu", getMenu)
	engine.POST("/updateUserPass", updateUserPass)
	engine.GET("/queryConsume", queryConsume)
	engine.GET("/queryCharge", queryCharge)
	engine.GET("/cashierLogin", cashierLogin)
	engine.POST("/charge", charge)
	engine.GET("customerServerLogin", customerServerLogin)
	engine.POST("/service", service)
	engine.GET("/validUserId", validUserId)
	engine.POST("/changeMenu", changeMenu)
	engine.POST("/consume", consume)
	engine.GET("/serviceStatistics", serviceStatistics)
	engine.GET("/adminLogin", adminLogin)
	engine.POST("/resetPassword", resetPassword)

	// 监听运行
	err = engine.Run(":7384")
	if err != nil {
		panic(err)
	}
}
