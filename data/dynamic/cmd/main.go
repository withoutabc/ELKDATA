package main

import (
	"ELKDATA/data/dynamic/api"
	"ELKDATA/data/dynamic/initialize"
)

func main() {
	initialize.InitShanghaiTime()
	initialize.InitLogger()
	api.InitRouters()
}
