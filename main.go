package main

import (
	"application/middleware"
	"application/router/api"
	"application/tools"

	"gitee.com/frankyu365/gocrypto/ecc"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := tools.ParseConfig("./config/config.json")
	if err != nil {
		panic(err)
	}
	//sdkuse初始化sdk 组织等信息
	tools.SdkSetup()
	ecc.GenerateECCKey(256, "./kgc/register/")
	ecc.GenerateECCKey(256, "./kgc/login/")
	app := gin.Default()
	app.Use(middleware.CORSMiddleware())
	api.InitAPI(app)
	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}
