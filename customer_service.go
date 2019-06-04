package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ServiceStatistics struct {
	Uid   string `json:"ser_id" gorm:"column:uid"`
	Count int32  `json:"cnt" gorm:"column:cnt"`
	Name  string `json:"name" gorm:"column:name"`
}

// 客服服务次数统计
func serviceStatistics(c *gin.Context) {
	// 和客户的差不多吧
	// 按照月份统计好了
	var res = gin.H{
		"success": false,
		"err":     "",
		"records": []ServiceStatistics{},
	}
	//idx, err := strconv.ParseInt(c.Query("index"), 10, 31)
	//if err != nil {
	//	res["err"] = err.Error()
	//	c.JSON(http.StatusOK, res)
	//	return
	//}
	year, err := strconv.ParseInt(c.DefaultQuery("year", "0"), 10, 31)
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	month, err := strconv.ParseInt(c.DefaultQuery("month", "0"), 10, 31)
	//_ = year
	//_ = month

	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	// 开始统计
	// 客服服务次数统计
	// 按照月份统计吧
	var recs []ServiceStatistics
	err = db.Raw("SELECT uid, name, (SELECT COUNT(*) FROM services WHERE ser_id = uid AND year = ? AND month = ?) AS cnt FROM cus_serv;", year, month).Find(&recs).Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	//fmt.Println(recs)
	res["records"] = recs
	res["success"] = true
	c.JSON(http.StatusOK, res)

}
