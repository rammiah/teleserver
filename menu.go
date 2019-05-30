package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMenu(c *gin.Context) {
	var menus = make([]Menu, 0)
	var res = gin.H{
		"status":  http.StatusOK,
		"err":     "",
		"records": menus,
	}
	err := db.Find(&menus).Error

	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res["records"] = menus
	c.JSON(http.StatusOK, res)
	fmt.Println(res)
}
