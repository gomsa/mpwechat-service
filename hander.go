package main

import (
	"context"
	"strings"

	authClient "github.com/gomsa/auth-service/client"
	authPd "github.com/gomsa/auth-service/proto/auth"
	userClient "github.com/gomsa/user-service/client"
	userPd "github.com/gomsa/user-service/proto/user"

	mp "github.com/gomsa/mpwechat-service/proto/mpwechat"
	"github.com/gomsa/mpwechat-service/service"
)

type hander struct {
	repo         service.Repository
	oauthService service.Oauth
}

func (srv *hander) Auth(ctx context.Context, req *mp.Request, res *mp.Token) (err error) {
	session, err := srv.oauthService.Session(req.Code)
	if err != nil {
		return err
	}
	user := &mp.User{}
	user, err = srv.repo.GetByOpenid(session.OpenId)
	// 如果不存在用户先创建用户
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			// 创建新用户
			// bug 无用户名创建用户可能引起 bug
			resp, err := userClient.Users.Create(context.TODO(), &userPd.User{
				Origin: serviceName,
			})
			if err != nil {
				return err
			}
			user = &mp.User{
				Id:         resp.User.Id,
				Openid:     session.OpenId,
				Unionid:    session.UnionId,
				SessionKey: session.SessionKey,
			}
			// 写入微信用户信息
			err = srv.repo.Create(user)
			if err != nil {
				return err
			}
		}
	}
	token, err := authClient.Auth.AuthById(
		context.TODO(),
		&authPd.User{
			Id: user.Id,
		},
	)
	if err != nil {
		return err
	}
	res.Token = token.Token
	return nil
}
func (srv *hander) UserInfo(ctx context.Context, req *mp.Request, res *mp.User) (err error) {
	return nil
}
