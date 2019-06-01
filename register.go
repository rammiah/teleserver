package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func userRegister(c *gin.Context) {
	// 用户注册接口
	var data = struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
		Menu int32  `json:"menu"`
	}{}
	err := c.ShouldBindJSON(&data)

	var res = gin.H{
		"success": false,
		"err":     "",
		"uid":     "",
	}
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	var user User
	user.Pass = data.Pass
	user.Name = data.Name
	user.Mid = data.Menu
	user.Money = 0
	user.Overdue = false
	// 获取编号
	var uid = struct {
		Uid string
	}{}
	tx := db.Table(user.TableName()).Begin()
	err = tx.Table(user.TableName()).Select("uid").Order("uid DESC").First(&uid).Error
	if err == gorm.ErrRecordNotFound {
		uid.Uid = "000000"
	} else if err == nil {
		num, _ := strconv.ParseInt(uid.Uid, 10, 0)
		// 获取uid，并发如何控制？
		uid.Uid = fmt.Sprintf("%06d", num+1)
	} else {
		tx.Rollback()
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	user.Uid = uid.Uid
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	tx.Commit()
	res["success"] = true
	res["uid"] = uid.Uid
	c.JSON(http.StatusOK, res)
}
