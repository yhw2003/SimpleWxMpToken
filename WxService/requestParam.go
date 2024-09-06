package wxservice

type TokenRequestTokenParam struct {
	GrantType string `json:"grant_type"`
	AppId     string `json:"appid"`
	Secret    string `json:"secret"`
}
