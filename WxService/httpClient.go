package wxservice

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var httpClient *http.Client

var appConfig *config

type config struct {
	grantType string
	appId     string
	appSecret string
}

func Init() {
	httpClient = &http.Client{}
	appConfig = &config{}
	appId := os.Getenv("APPID")
	appSecret := os.Getenv("APPSECRET")
	grantType := "client_credential"
	log.Printf("初始化WxService: appId: %s, appSecret: %s, grantType: %s\n", appId, appSecret, grantType)
	appConfig.appId = appId
	appConfig.appSecret = appSecret
	appConfig.grantType = grantType
}

func GetNewToken() *TokenReturnParse {
	return getNewToken(TokenRequestTokenParam{
		GrantType: appConfig.grantType,
		AppId:     appConfig.appId,
		Secret:    appConfig.appSecret,
	})
}

func getNewToken(param TokenRequestTokenParam) *TokenReturnParse {
	resp, err := httpClient.Get(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s", param.GrantType, param.AppId, param.Secret))
	if err != nil {
		log.Printf("获取token失败: %v\n", err)
		return nil
	}
	defer resp.Body.Close()
	tokenReturnParse := &TokenReturnParse{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取token返回body失败: %v\n", err)
		return nil
	}
	var errorParse = &struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}{}
	err = json.Unmarshal(body, errorParse)
	if err != nil {
		log.Printf("解析token返回body失败: %v\n", err)
		return nil
	}
	if errorParse.ErrCode != 0 {
		log.Printf("获取token失败: errcode: %d, errmsg: %s\n", errorParse.ErrCode, errorParse.ErrMsg)
		return nil
	}
	json.Unmarshal(body, tokenReturnParse)
	log.Printf("获取token成功: %s\n", tokenReturnParse.AccessToken)
	return tokenReturnParse
}
