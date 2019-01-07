package main

import (
	"fmt"
	"log"

	mp "github.com/gomsa/mpwechat-service/proto/wechat"
	local "github.com/gomsa/mpwechat-service/service"
	"github.com/jinzhu/gorm"
	micro "github.com/micro/go-micro"
)

func autoMigrate(db *gorm.DB) {
	user := &mp.User{}
	if !db.HasTable(&user) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&user).
			AddUniqueIndex("idx_user_openid", "openid")
	}
}
func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	// 自动数据库迁移
	autoMigrate(db)
	srv := micro.NewService(
		micro.Name("gomsa.mpwechat"),
		micro.Version("latest"),
	)
	srv.Init()
	// 用户仓库 db 接口实现
	repo := &local.UserRepository{db}
	// oauth 微信小程序服务
	oauth := &local.OauthService{}
	// Register handler
	mp.RegisterMpWechatHandler(srv.Server(), &service{repo, oauth})
	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
	log.Println("serviser run ...")
}
