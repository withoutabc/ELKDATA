package api

import (
	"ELKDATA/data/dynamic/initialize"
	"ELKDATA/data/dynamic/middleware"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	html = "visit.html"
)

func InitRouters() {
	r := gin.Default()
	r.GET("/visit", middleware.ReqTimeInfo(), Visit)
	r.GET("/ip", GetIp)
	r.Run(":5888")
}

func Visit(c *gin.Context) {
	c.File("./visit.html")
}

func GetIp(c *gin.Context) {
	ip := c.ClientIP()
	ipdb := initialize.InitIp()
	results, err := ipdb.Get_all(ip)
	if err != nil {
		hlog.Error("get ip_info failed,", err)
	}
	hlog.Infof("country:%s,region:%s,city:%s", results.Country_long, results.Region, results.City)
	c.JSON(http.StatusOK, IP{
		Country: results.Country_long,
		Region:  results.Region,
		City:    results.City,
	})
}

type IP struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}
