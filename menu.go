package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMenu(c *gin.Context) {
	var menus []Menu
	err := db.Find(&menus).Error

	var res = gin.H{
		"status":  http.StatusOK,
		"err":     "",
		"records": "[]",
	}

	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	res["records"] = menus
	c.JSON(http.StatusOK, res)
	fmt.Println(res)
}
