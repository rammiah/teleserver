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
