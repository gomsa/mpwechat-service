package service

import (
	"github.com/chanxuehong/wechat/mp/oauth2"
)

var (
	oauth2Endpoint *oauth2.Endpoint = oauth2.NewEndpoint(appID, appSecret)
)

// Oauth 微信小程序认证接口
type Oauth interface {
	Session(code string) (session *oauth2.Session, err error)
	UserInfo(openid string) (sessionInfo *oauth2.SessionInfo, err error)
}

// OauthService 微信小程序认证服务实现
type OauthService struct {
}

// Session 获取微信小程序登录 session
func (srv *OauthService) Session(code string) (session *oauth2.Session, err error) {
	session, err = oauth2.GetSession(oauth2Endpoint, code)
	if err != nil {
		return session, err
	}
	return session, err
}

// UserInfo 获取微信小程序用户信息
func (srv *OauthService) UserInfo(openid string) (sessionInfo *oauth2.SessionInfo, err error) {

	return sessionInfo, err
}
