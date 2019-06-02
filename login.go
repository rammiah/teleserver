package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户登录
func userLogin(c *gin.Context) {
	uid := c.Query("uid")
	pass := c.Query("pass")
	//fmt.Printf("uid: %q, pass: %q\n", uid, pass)
	var user User
	err := db.Table(user.TableName()).Where("uid = ?", uid).First(&user).Error
	//fmt.Println(user)
	var res = gin.H{
		"success": false,
		"name":    "",
		"err":     "",
		"money":   0,
	}
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if user.Pass == pass {
		res["name"] = user.Name
		res["success"] = true
		res["money"] = user.Money
		c.JSON(http.StatusOK, res)
	} else {
		res["err"] = "wrong number or password"
		c.JSON(http.StatusOK, res)
	}
}

// 客服登录
func customerServerLogin(c *gin.Context) {
	uid := c.Query("uid")
	pass := c.Query("pass")
	//fmt.Printf("uid: %q, pass: %q\n", uid, pass)
	var user CustomerService
	err := db.Table(user.TableName()).Where("uid = ?", uid).First(&user).Error
	//fmt.Println(user)
	var res = gin.H{
		"success": false,
		"name":    "",
		"err":     "",
	}
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if user.Pass == pass {
		res["name"] = user.Name
		res["success"] = true
		c.JSON(http.StatusOK, res)
	} else {
		res["err"] = "wrong number or password"
		c.JSON(http.StatusOK, res)
	}
}

// 收款员登录
func cashierLogin(c *gin.Context) {
	uid := c.Query("uid")
	pass := c.Query("pass")
	//fmt.Printf("uid: %q, pass: %q\n", uid, pass)
	var user Cashier
	err := db.Table(user.TableName()).Where("uid = ?", uid).First(&user).Error
	var res = gin.H{
		"success": false,
		"name":    "",
		"err":     "",
	}
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if user.Pass == pass {
		res["name"] = user.Name
		res["success"] = true
		c.JSON(http.StatusOK, res)
	} else {
		res["err"] = "wrong number or password"
		c.JSON(http.StatusOK, res)
	}
}

// 管理员登录
func adminLogin(c *gin.Context) {
	uid := c.Query("uid")
	pass := c.Query("pass")
	//fmt.Printf("uid: %q, pass: %q\n", uid, pass)
	var user Admin
	err := db.Table(user.TableName()).Where("uid = ?", uid).First(&user).Error
	//fmt.Println(user)
	var res = gin.H{
		"success": false,
		"name":    "",
		"err":     "",
	}
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	if user.Pass == pass {
		res["name"] = user.Name
		res["success"] = true
		c.JSON(http.StatusOK, res)
	} else {
		res["err"] = "wrong number or password"
		c.JSON(http.StatusOK, res)
	}
}
