package initialize

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gin-gonic/gin"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"time"
)

const (
	LogFilePath = "./tmp/"
)

var Loc *time.Location

func InitShanghaiTime() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	Loc = loc
}

func InitHLogger() {
	// Customizable output directory.
	logFilePath := LogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().In(Loc).Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := hertzlogrus.NewLogger()
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	logger.SetOutput(lumberjackLogger)

	logger.SetLevel(hlog.LevelDebug)

	hlog.SetLogger(logger)
}

func LoggerWithFormatter(params gin.LogFormatterParams) string {
	return fmt.Sprintf(
		"timestamp:%s,status_code:%d,client_ip:%s,latency:%s,method:%s,path:%s\n",
		time.Now().In(Loc).Format("2006-01-02 15:04:05"),
		params.StatusCode, // 状态码
		params.ClientIP,   // 客户端ip
		params.Latency,    // 请求耗时
		params.Method,     // 请求方法
		params.Path,       // 路径
	)
}

func InitGINLogger() *os.File {
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
	return f
}
