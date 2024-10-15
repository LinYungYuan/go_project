package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"go_project/internal/app"
	"go_project/internal/config"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// 定義一個命令行標誌來指定環境
	env := flag.String("env", "dev", "Specify the environment (dev, uat, prd)")
	flag.Parse()

	// 根據命令行參數設置環境變量
	os.Setenv("ENVIRONMENT", *env)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	var name string = "123"
	fmt.Println(name)

	app, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}
	// 添加 Swagger 路由
	app.Router().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := app.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
