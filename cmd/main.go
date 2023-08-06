package main

import (
	"n8n-workflow/config"
	"n8n-workflow/routes"
)

func main() {
	config.InitConfig()
	ginRouter := routes.NewRouter()
	ginRouter.Run(config.Conf.Server.Host + ":" + config.Conf.Server.Port)
}
