package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func service(c *gin.Context) {
	// 这个一般是客服接电话结束后会有一个记录
	// 但是这个系统里并没有设计接受请求这个东西
	// 所以直接由客服发送记录好了
	//fmt.Println("add service record")
	var data = struct {
		CustomerServiceId string `json:"customer_service_id" binding:"required"`
		UserId            string `json:"user_id" binding:"required"`
	}{}

	var res = gin.H{
		"success": false,
		"err":     "",
	}

	err := c.ShouldBindJSON(&data)
	//fmt.Println(data)
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	// 准备创建记录
	var rec ServiceRecord
	rec.UserId = data.UserId
	rec.SerId = data.CustomerServiceId
	// 填充时间部分
	now := time.Now()
	rec.Year = int32(now.Year())
	rec.Month = int32(now.Month())
	rec.Day = int32(now.Day())
	rec.Tm = []byte(now.Format("15:04:05"))
	tx := db.Begin()
	err = tx.Save(&rec).Error
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}
	// id
	//fmt.Println(rec.Id)
	//fmt.Printf("%q\n", res)
	res["success"] = true
	tx.Commit()
	//fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

func validUserId(c *gin.Context) {
	var uid = c.Query("uid")
	var count = 0
	var res = gin.H{
		"success": false,
		"err":     "",
		"valid":   false,
	}
	err := db.Table("user").Where("uid = ?", uid).Count(&count).Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res["valid"] = count > 0
	res["success"] = true
	c.JSON(http.StatusOK, res)
}
