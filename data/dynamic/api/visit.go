package api

import (
	"ELKDATA/data/dynamic/initialize"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

const (
	VisitHtml = "./html/visit.html"
	SlowHtml  = "./html/slow.html"
)

func InitRouters() {
	r := gin.Default()
	logFilePath := "./tmp/"
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r.Use(gin.LoggerWithFormatter(initialize.LoggerWithFormatter))

	r.GET("/visit", Visit)
	r.GET("/ip", GetIp)

	r.GET("/slow", Slow)
	r.Run(":5888")
}

func Slow(c *gin.Context) {
	time.Sleep(time.Millisecond * 500)
	c.File(SlowHtml)
}

func Visit(c *gin.Context) {
	c.File(VisitHtml)
}

func GetIp(c *gin.Context) {
	ip := c.ClientIP()
	ipdb := initialize.InitIp()
	results, err := ipdb.Get_all(ip)
	if err != nil {
		hlog.Error("get ip_info failed,", err)
	}
	if results.Region == "Hong Kong" || results.Region == "Tai Wan" || results.Region == "Macao" || results.Region == "Taiwan" {
		results.Country_long = "China"
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
