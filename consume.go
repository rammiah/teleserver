package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func consume(c *gin.Context) {
	// 消费记录添加
	var data = struct {
		Uid  string  `json:"uid" binding:"required"`
		Cost float32 `json:"cost" binding:"required"`
	}{}

	var res = gin.H{
		"success": false,
		"err":     "",
	}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	// 信息获取成功
	var consume Consume
	consume.Cost = data.Cost
	consume.UserId = data.Uid
	now := time.Now()
	consume.Year = int32(now.Year())
	consume.Month = int32(now.Month())
	consume.Day = int32(now.Day())
	consume.Tm = []byte(now.Format("15:04:05"))
	tx := db.Begin()
	err = tx.Table(consume.TableName()).
		Save(&consume).
		Error
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}
	var user User
	err = tx.Table(user.TableName()).
		Where("uid = ?", data.Uid).
		First(&user).
		Error
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}

	if user.Overdue {
		res["err"] = "user is overdue"
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}
	user.Money -= data.Cost
	user.Overdue = user.Money < 0
	err = tx.Save(&user).
		Error
	if err != nil {
		res["err"] = err.Error()
		tx.Rollback()
		c.JSON(http.StatusOK, res)
		return
	}
	// 完成
	tx.Commit()
	res["success"] = true
	c.JSON(http.StatusOK, res)
}
