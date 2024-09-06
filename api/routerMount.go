package api

import (
	v1 "wx_token_service/api/v1"

	"github.com/gin-gonic/gin"
)

func MountRouter(engine *gin.Engine) {
	engine.GET("/token/:access_token", v1.GetTokenHandler)
}

func Init() {
	v1.Init()
}
