package api

import (
	"ELKDATA/data/dynamic/initialize"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
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
	url := fmt.Sprintf("http://ip-api.com/json/%s?fields=61439&lang=zh-CN", ip)

	resp, err := http.Get(url)
	if err != nil {
		hlog.Error("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		hlog.Error("读取响应失败:", err)
		return
	}

	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		hlog.Error("解析JSON失败:", err)
		return
	}
	if err != nil {
		hlog.Error("get ip_info failed:", err)
	}
	if ipInfo.Country == "香港" {
		ipInfo.Country = "中国香港"
	}
	if ipInfo.RegionName == "台湾" {
		ipInfo.Country = "中国台湾"
	}
	if ipInfo.City == "澳门" {
		ipInfo.Country = "中国澳门"
	}
	hlog.Infof("country:%s,region:%s,city:%s,latitude:%f,longitude:%f", ipInfo.Country, ipInfo.RegionName, ipInfo.City, ipInfo.Lat, ipInfo.Lon)
	c.JSON(http.StatusOK, IP{
		Country: ipInfo.Country,
		Region:  ipInfo.RegionName,
		City:    ipInfo.City,
	})
}

type IP struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}

type IPInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}
