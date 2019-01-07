package cli

import (
	pb "github.com/gomsa/user-service/proto/auth"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

// Auth 用户客户端
var Auth pb.AuthClient

func init() {
	cmd.Init()
	// 创建 user-service 微服务的客户端
	Auth = pb.NewAuthClient("gomsa.auth", microclient.DefaultClient)
}
