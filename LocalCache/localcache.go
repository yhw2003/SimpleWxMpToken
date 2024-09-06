package localcache

import (
	"log"
	"sync"
	"time"
	wxservice "wx_token_service/WxService"
)

type LocalCache struct {
	wxToken     string
	expiresTime time.Time
}

var localCache *LocalCache

var cacheLock = &sync.Mutex{}

func Init() {
	localCache = &LocalCache{}
	tokenParse := wxservice.GetNewToken()
	localCache.wxToken = tokenParse.AccessToken
	localCache.expiresTime = time.Now().Add(time.Duration(tokenParse.ExpiresIn-120) * time.Second)
}

// GetCacheToken 获取缓存的token
// 如果token过期，会重新获取
func GetCacheToken() string {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	if time.Now().After(localCache.expiresTime) {
		tokenParse := wxservice.GetNewToken()
		localCache.wxToken = tokenParse.AccessToken
		localCache.expiresTime = time.Now().Add(time.Duration(tokenParse.ExpiresIn-120) * time.Second)
		log.Println("尝试更新token缓存")
	}
	return localCache.wxToken
}
