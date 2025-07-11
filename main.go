package main

import (
	"monit/config"
	"monit/route"
)

func main() {
	config.InitConfig()
	r := route.SetupRouter()
	r.Run(config.AppConfig.App.Port)
}
