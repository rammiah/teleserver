package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 测试服务器是否可用
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
