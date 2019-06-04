package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMenu(c *gin.Context) {
	var menus []Menu
	var res = gin.H{
		"success": false,
		"err":     "",
		"records": menus,
	}
	err := db.Find(&menus).Error

	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res["success"] = true
	res["records"] = menus
	c.JSON(http.StatusOK, res)
}

func changeMenu(c *gin.Context) {
	var data = struct {
		Uid     string `json:"uid" binding:"required"`
		NewMenu int32  `json:"new_menu"`
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
	// 更新套餐
	var user User
	err = db.Table(user.TableName()).
		Where("uid = ?", data.Uid).
		First(&user).
		Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	user.Mid = data.NewMenu
	err = db.Table(user.TableName()).
		UpdateColumn("mid").
		Save(&user).
		Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res["success"] = true
	c.JSON(http.StatusOK, res)
}
