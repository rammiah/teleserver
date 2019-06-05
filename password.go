package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func updateUserPass(c *gin.Context) {
	var res = gin.H{
		"success": false,
		"err":     "",
	}

	var data = struct {
		Uid     string `json:"uid" binding:"required"`
		OldPass string `json:"old_pass" binding:"required"`
		NewPass string `json:"new_pass" binding:"required"`
	}{}
	//fmt.Println("bind json")
	err := c.ShouldBindJSON(&data)

	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	var user User
	fmt.Println("find user")
	fmt.Println(data)
	err = db.Where("uid = ? AND pass = ?", data.Uid, data.OldPass).First(&user).Error
	if err != nil {
		//fmt.Println(err)
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	//fmt.Println(user)

	user.Pass = data.NewPass
	err = db.Table(user.TableName()).Where("uid = ?", user.Uid).UpdateColumn("pass").Save(&user).Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res["success"] = true
	c.JSON(http.StatusOK, res)
}

const (
	_User = iota
	_CustomerService
	_Cashier
)

func resetPassword(c *gin.Context) {
	var data = struct {
		Type int32  `json:"type"`
		Uid  string `json:"uid" binding:"required"`
		Pass string `json:"pass" binding:"required"`
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
	fmt.Println(data)
	// 根据情况查找
	switch data.Type {
	case _User:
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
		user.Pass = data.Pass
		err = db.Save(&user).Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		res["success"] = true
		c.JSON(http.StatusOK, res)
	case _CustomerService:
		var cs CustomerService
		err = db.Table(cs.TableName()).
			Where("uid = ?", data.Uid).
			First(&cs).
			Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		cs.Pass = data.Pass
		err = db.Save(&cs).Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		res["success"] = true
		c.JSON(http.StatusOK, res)
	case _Cashier:
		var cashier Cashier
		err = db.Table(cashier.TableName()).
			Where("uid = ?", data.Uid).
			First(&cashier).
			Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		cashier.Pass = data.Pass
		err = db.Save(&cashier).Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		res["success"] = true
		c.JSON(http.StatusOK, res)
	default:
		res["err"] = "no such type"
		c.JSON(http.StatusOK, res)
	}
}
