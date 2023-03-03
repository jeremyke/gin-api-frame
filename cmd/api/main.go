package main

import (
	"gin-api-frame/app/global/variable"
	"gin-api-frame/bootstrap"
	"gin-api-frame/routers"
)

func main() {
	bootstrap.InitServer()
	r := routers.InitApiRouter(false)
	r.Run(":" + variable.HttpPort)
}
