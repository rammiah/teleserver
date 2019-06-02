package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	_Month = iota
	_Day
	_Once
)

type MonthRecord struct {
	// 按照月收的没有流水号了
	Year  int32   `json:"year" `
	Month int32   `json:"month"`
	Cost  float32 `json:"cost"`
}

type DayRecord struct {
	Year  int32   `json:"year"`
	Month int32   `json:"month"`
	Day   int32   `json:"day"`
	Cost  float32 `json:"cost"`
}

type OnceRecord struct {
	Year  int32   `json:"year"`
	Month int32   `json:"month"`
	Day   int32   `json:"day"`
	Tm    []byte  `json:"tm"`
	Cost  float32 `json:"cost"`
}

func queryConsume(c *gin.Context) {
	var res = gin.H{
		"success": false,
		"left":    0,
		"err":     "",
		"records": nil,
	}

	var idx, err = strconv.ParseInt(c.DefaultQuery("index", "0"), 10, 32)
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	//fmt.Println(idx)
	queryType, err := strconv.ParseInt(c.DefaultQuery("type", "0"), 10, 32)
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	//fmt.Println(queryType)
	//c.String(http.StatusOK, "hello")
	var uid = c.DefaultQuery("uid", "")
	//fmt.Println(uid)

	switch queryType {
	case _Month:
		// 按照每月为单位
		var count = 0
		err = db.Table("consume").
			Select("year, month").
			Where("uid = ?", uid).
			Group("year, month").
			Count(&count).
			Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		//fmt.Println(count)
		var recs []MonthRecord

		err = db.Table("consume").
			Select("year, month, SUM(cost) AS cost").
			Where("uid = ?", uid).
			Group("year, month").
			//Group("month").
			Order("year DESC, month DESC").
			Limit(10).
			Offset(idx).
			Find(&recs).
			Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		// 计算剩余数据条数
		if count-int(idx)-10 < 0 {
			res["left"] = 0
		} else {
			res["left"] = count - int(idx) - 10
		}
		//fmt.Printf("%+v\n", recs)
		res["success"] = true
		res["records"] = recs
		c.JSON(http.StatusOK, res)
		return
	case _Day:
		var count = 0
		err = db.Table("consume").
			Select("year, month，day").
			Where("uid = ?", uid).
			Group("year, month, day").
			Count(&count).
			Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		//fmt.Println(count)
		var recs []DayRecord
		err = db.Table("consume").
			Select("year, month, day, SUM(cost) AS cost").
			Where("uid = ?", uid).
			Group("year, month, day").
			Order("year DESC, month DESC, day DESC").
			Limit(10).
			Offset(idx).
			Find(&recs).
			Error
		// 计算剩余数据条数
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		if count-int(idx)-10 < 0 {
			res["left"] = 0
		} else {
			res["left"] = count - int(idx) - 10
		}
		//fmt.Printf("%+v\n", recs)
		res["success"] = true
		res["records"] = recs
		c.JSON(http.StatusOK, res)
		return
	case _Once:
		var count = 0
		err = db.Table("consume").
			Where("uid = ?", uid).
			Count(&count).
			Error
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		//fmt.Println(count)
		var recs []OnceRecord
		err = db.Table("consume").
			Select("year, month, day, tm, cost").
			Where("uid = ?", uid).
			Order("year DESC, month DESC, day DESC, tm DESC").
			Limit(10).
			Offset(idx).
			Find(&recs).
			Error
		// 计算剩余数据条数
		if err != nil {
			res["err"] = err.Error()
			c.JSON(http.StatusOK, res)
			return
		}
		if count-int(idx)-10 < 0 {
			res["left"] = 0
		} else {
			res["left"] = count - int(idx) - 10
		}

		//fmt.Printf("%+v\n", recs)
		res["success"] = true
		res["records"] = recs
		c.JSON(http.StatusOK, res)
		return
	default:
		res["err"] = "unknown query type"
		c.JSON(http.StatusOK, res)
		return
	}
}

func queryCharge(c *gin.Context) {
	// 查询缴费情况
	var uid = c.DefaultQuery("uid", "")
	var idx, err = strconv.ParseInt(c.DefaultQuery("index", "0"), 10, 32)
	var res = gin.H{
		"success": false,
		"left":    0,
		"err":     "",
		"records": []Charge{},
	}

	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	var count = 0
	// 获取收费表中此人记录数目
	err = db.Table("charge").Where("user_id = ?", uid).Count(&count).Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}
	var charges []Charge
	err = db.Table("charge").
		Where("user_id = ?", uid).
		Order("year DESC").
		Order("month DESC").
		Order("day  DESC").
		Order("tm DESC").
		Limit(10).
		Offset(idx).
		Find(&charges).Error
	if err != nil {
		res["err"] = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	res["success"] = true
	// 将结果的时间从base64转回字符串数据
	res["records"] = charges
	if count-int(idx)-10 < 0 {
		res["left"] = 0
	} else {
		res["left"] = count - int(idx) - 10
	}

	c.JSON(http.StatusOK, res)
}
