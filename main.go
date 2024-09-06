package main

import (
	"log"
	"os"
	localcache "wx_token_service/LocalCache"
	wxservice "wx_token_service/WxService"
	"wx_token_service/api"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("[WxTokenService]")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("读取环境变量失败: %v\n", err)
	}
	log.Println("读取环境变量成功")
	wxservice.Init()
	localcache.Init()
	api.Init()
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
		log.Println("设置gin为release模式")
	}
}

func main() {
	engine := gin.Default()
	api.MountRouter(engine)
	engine.Run(":8080")
}
