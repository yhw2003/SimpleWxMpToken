# 微信小程序token缓存服务
这是一个非常简单的组件，用于在微服务中统一获取微信小程序的access_token，并缓存在内存中，避免多个服务竞争token。
用法:
```bash
cp .env.simple .env
```
在内填写对用数据,```SERVICE_TOKEN```是服务间鉴权token,另外两个从微信开放平台获取.