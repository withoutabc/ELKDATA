package initialize

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/ip2location/ip2location-go/v9"
)

const dbFile = "./../IP2LOCATION-LITE-DB11.BIN"

func InitIp() *ip2location.DB {
	ipDb, err := ip2location.OpenDB(dbFile)

	if err != nil {
		hlog.Fatal("initialize ipDb failed", err)
		return nil
	}
	return ipDb
}
