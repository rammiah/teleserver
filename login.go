package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// 用户登录
func userLogin(c *gin.Context) {
	uid := c.Query("uid")
	pass := c.Query("pass")
	fmt.Printf("uid: %q, pass: %q\n", uid, pass)
	var user User
	err := db.Table(user.TableName()).Where("uid = ?", uid).First(&user).Error
	fmt.Println(user)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"login":  false,
		})
		return
	} else if err != nil {
		panic(err)
	}

	if user.Pass == pass {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"login":  true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"login":  false,
		})
	}
}

// 客服登录
func customerServerLogin(name, pass string) {
	var con_serv CustomerServer
	err := db.Table(con_serv.TableName()).Where("uid = ?", name).First(&con_serv).Error
	if err == gorm.ErrRecordNotFound {
		//return false

	} else if err != nil {
		panic(err)
	}

	//return con_serv.Pass == pass
}

// 收款员登录
func cashierLogin(name, pass string) {
	var cashier Cashier
	err := db.Table(cashier.TableName()).Where("uid = ?", name).First(&cashier).Error

	if err == gorm.ErrRecordNotFound {
		//return false
	} else if err != nil {
		panic(err)
	}

	//return cashier.Pass == pass
}

// 管理员登录
func adminLogin(name, pass string) {
	var admin Admin
	err := db.Table(admin.TableName()).Where("uid = ?", name).First(&admin).Error

	if err != gorm.ErrRecordNotFound {
		//return false
	} else if err != nil {
		panic(err)
	}

	//return admin.Pass == pass
}
