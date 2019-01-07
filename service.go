package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/gomsa/user-service/proto/auth"

	client "github.com/gomsa/mpwechat-service/client"
	mp "github.com/gomsa/mpwechat-service/proto/wechat"
	ext "github.com/gomsa/mpwechat-service/service"
)

type service struct {
	repo         ext.Repository
	oauthService ext.Oauth
}

func (srv *service) Auth(ctx context.Context, req *mp.Request, res *mp.Token) (err error) {
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
			resp, err := client.Auth.Create(context.TODO(), &pb.User{
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
	token, err := client.Auth.AuthById(context.TODO(), &pb.User{
		Id: user.Id,
	})
	if err != nil {
		return err
	}
	res.Token = token.Token
	return nil
}
func (srv *service) UserInfo(ctx context.Context, req *mp.Request, res *mp.User) (err error) {

	log.Println(req)
	return nil
}
