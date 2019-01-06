package main

import (
	"fmt"
	"log"

	mp "github.com/gomsa/mpwechat-service/proto/wechat"
	local "github.com/gomsa/mpwechat-service/service"
	micro "github.com/micro/go-micro"
)

func main() {
	srv := micro.NewService(
		micro.Name("gomsa.mpwechat"),
		micro.Version("latest"),
	)
	srv.Init()
	// oauth2 微信小程序服务
	oauth2 := &local.Oauth2Service{}
	// Register handler
	mp.RegisterMpWechatHandler(srv.Server(), &service{oauth2})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	log.Println("serviser run ...")
}
