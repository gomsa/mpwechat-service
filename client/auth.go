package cli

import (
	"os"

	pb "github.com/gomsa/user-service/proto/auth"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

var (
	// Auth 用户客户端
	Auth pb.AuthClient
	// authService 用户认证服务名称
	authService = os.Getenv("AUTH_SERVICE")
)

func init() {
	cmd.Init()
	// 创建 user-service 微服务的客户端
	Auth = pb.NewAuthClient(authService, microclient.DefaultClient)
}
