package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/gomsa/user-service/providers/database"

	auth "github.com/gomsa/mpwechat-service/auth"
	mp "github.com/gomsa/mpwechat-service/proto/mpwechat"
	"github.com/gomsa/mpwechat-service/service"

	micro "github.com/micro/go-micro"
)

var (
	// serviceName 服务名称
	serviceName = os.Getenv("SERVICE_NAME")
)

// func autoMigrate(db *gorm.DB) {
// 	user := &mp.User{}
// 	if !db.HasTable(&user) {
// 		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user).
// 			AddUniqueIndex("idx_user_openid", "openid")
// 	}
// }
func main() {
	// db, err := CreateConnection()
	// defer db.Close()

	// if err != nil {
	// 	log.Fatalf("connect error: %v\n", err)
	// }
	// // 自动数据库迁移
	// autoMigrate(db)
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.WrapHandler(auth.Wrapper),
	)
	srv.Init()
	// 用户仓库 db 接口实现
	repo := &service.UserRepository{db.DB}
	// oauth 微信小程序服务
	oauth := &service.OauthService{}
	// Register handler
	mp.RegisterMpWechatHandler(srv.Server(), &hander{repo, oauth})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	log.Println("serviser run ...")
}
