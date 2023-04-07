package main

import (
	biz_config "biz-c-service/biz-config"
	"biz-c-service/router"

	"github.com/isyscore/isc-gobase/server"
)

func main() {
	biz_config.InitConfig()
	router.Register()
	server.Run()
}
