package service

import "os"

var (
	appID         = os.Getenv("WECHAT_APP_ID")
	appSecret     = os.Getenv("WECHAT_APP_SECRET")
	oriID         = os.Getenv("WECHAT_ORIID")
	token         = os.Getenv("WECHAT_TOKEN")
	encodedAESKey = os.Getenv("WECHAT_ENCODE_AES_KEY")
)
