package main

import (
	"context"

	mp "github.com/gomsa/mpwechat-service/proto/wechat"
	ext "github.com/gomsa/mpwechat-service/service"
)

type service struct {
	mp ext.Oauth2
}

func (srv *service) Auth(ctx context.Context, req *mp.Request, res *mp.Token) (err error) {
	srv.mp.Session(req.Code)
	return nil
}
