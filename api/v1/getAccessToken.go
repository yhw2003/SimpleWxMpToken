package v1

import (
	"log"
	"net/http"
	"os"
	localcache "wx_token_service/LocalCache"

	"github.com/gin-gonic/gin"
)

var appServiceToken string

func Init() {
	appServiceToken = os.Getenv("SERVICE_TOKEN")
	log.Printf("初始化GetAccessToken: service_token: %s\n", appServiceToken)
}

func GetTokenHandler(c *gin.Context) {
	serviceToken := c.Param("access_token")
	if serviceToken != appServiceToken {
		log.Printf("service_token不正确, got: %s, want: %s", serviceToken, appServiceToken)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "service_token not correct",
		})
		return
	}
	token := localcache.GetCacheToken()
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
