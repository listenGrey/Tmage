package main

import (
	"Tmage/controller"
	"Tmage/logger"
	"Tmage/routers"
	"Tmage/setting"
	"fmt"
)

// @title Tmage
// @version 1.0
// @description Tmage backend API testing.

// @host 127.0.0.1:8088
// @BasePath /api/v1/
func main() {
	if err := setting.Init(); err != nil {
		fmt.Printf("配置文件加载错误：%s\n", err)
		return
	}

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("logger初始化错误：%s\n", err)
		return
	}

	if err := controller.InitTranslator("zh"); err != nil {
		fmt.Printf("validator 初始化失败：%s\n", err)
		return
	}

	r := routers.SetupRouter(setting.Conf.Mode)
	r.Run(":8088")
	err := r.Run(fmt.Sprintf(""))
	if err != nil {
		fmt.Printf("Run server failed, err: %s\n", err)
		return
	}
}
