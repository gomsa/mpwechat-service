package service

import (
	"log"
	"os"

	mp "github.com/gomsa/mpwechat-service/proto/wechat"
)

var (
	appID         = os.Getenv("APP_KEY")
	appSecret     = os.Getenv("APP_KEY")
	oriID         = os.Getenv("APP_KEY")
	token         = os.Getenv("APP_KEY")
	encodedAESKey = os.Getenv("APP_KEY")
)

// Oauth2 微信小程序认证接口
type Oauth2 interface {
	Session(code string) (token *mp.Token, err error)
}

// Oauth2Service 微信小程序认证服务实现
type Oauth2Service struct {
}

// Session 获取微信小程序登录 session
func (srv *Oauth2Service) Session(code string) (token *mp.Token, err error) {
	log.Println(code)
	return nil, nil
}
