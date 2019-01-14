package auth

import (
	"context"
	"errors"
	"log"
	"strings"

	authClient "github.com/gomsa/auth-service/client"
	authPb "github.com/gomsa/auth-service/proto/auth"
	"github.com/gomsa/mpwechat-service/providers/config"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)

// ValidateMoethod 返回认证方式
func ValidateMoethod(method string) (action string) {
	auth := config.Conf.Validate["auth"]
	permission := config.Conf.Validate["permission"]
	if isSlice(auth, method) {
		return "auth"
	}
	if isSlice(permission, method) {
		return "permission"
	}
	return action
}

// isSlice 是否存在
func isSlice(slice []config.Spec, method string) bool {
	for _, value := range slice {
		if value.Label == method {
			return true
		}
	}
	return false
}

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
		// 三种方式
		// 1、auth 只验证登录状态
		// 2、permission 验证登录和访问 method 权限状态
		// 3、无状态任何请求都可以访问
		switch ValidateMoethod(req.Method()) {
		case "auth":
			// Auth here
			authResp, err := authClient.Auth.ValidateToken(context.Background(), &authPb.Request{
				Token: token,
			})
			log.Println("Auth Resp:", authResp)
			if err != nil {
				return err
			}
		case "permission":
			// Auth here
			authResp, err := authClient.Auth.ValidatePermission(context.Background(), &authPb.Request{
				Token:   token,
				Service: req.Service(),
				Method:  req.Method(),
			})
			log.Println("Auth Resp:", authResp)
			if err != nil {
				return err
			}
		default:
			log.Println("make dev")
			err = fn(ctx, req, resp)
		}
		err = fn(ctx, req, resp)
		return err
	}
}
