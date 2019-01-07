package auth

import (
	"context"
	"errors"
	"log"
	"strings"

	pb "github.com/gomsa/user-service/proto/auth"

	client "github.com/gomsa/mpwechat-service/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

// Wrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func Wrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := strings.Split(meta["Authorization"], "Bearer ")[1]

		// Auth here
		authResp, err := client.Auth.ValidateToken(context.Background(), &pb.Token{
			Token: token,
		})
		log.Println("Auth Resp:", authResp)
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
