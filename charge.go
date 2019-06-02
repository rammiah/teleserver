package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func charge(c *gin.Context) {
	// 首先获取参数
	var res = gin.H{
		"success": false,
		"err":     "",
	}
	//fmt.Println(res)

	var data = struct {
		CashierId string  `json:"cashier_id" binding:"required"`
		UserId    string  `json:"user_id" binding:"required"`
		Money     float32 `json:"money" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	//fmt.Printf("%q\n", data)
	//res["success"] = true
	//c.JSON(http.StatusOK, res)
	// 构建一个缴费记录
	var rec Charge
	rec.UserId = data.UserId
	rec.CashierId = data.CashierId
	rec.Money = data.Money
	// 还需设置年,月,日和时间
	now := time.Now()
	rec.Year = int32(now.Year())
	rec.Month = int32(now.Month())
	rec.Day = int32(now.Day())
	rec.Tm = []byte(now.Format("15:04:05"))
	tx := db.Begin()
	err = tx.Table("charge").Save(&rec).Error
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}

	// 给余额增加一下
	var user User
	err = tx.Table("user").
		Where("uid = ?", data.UserId).
		First(&user).
		Error
	//fmt.Println(user)
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}
	user.Money += data.Money
	user.Overdue = user.Money < 0
	//fmt.Println("save")
	err = tx.Save(user).
		Error
	//fmt.Println("after save")
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}
	//fmt.Println("send")
	res["success"] = true
	tx.Commit()
	c.JSON(http.StatusOK, res)
}
