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
