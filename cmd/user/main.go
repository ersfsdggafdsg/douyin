package main

import (
	"douyin/shared/middleware"
	"douyin/shared/initialize"
	user "douyin/shared/rpc/kitex_gen/user/userservice"
	"log"

	"douyin/cmd/user/pkg/model"
	"douyin/cmd/user/pkg/mysql"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("user.srv")
	svr := user.NewServer(&UserServiceImpl {
			Db: mysql.NewManager(initialize.InitMysql(
				"douyin", "zhihao", "douyin", &model.UserInfo{})),
		},
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
		server.WithRegistry(r),
		server.WithMiddleware(middleware.ShowCallingMiddleware),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
