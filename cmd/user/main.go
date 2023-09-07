package main

import (
	"douyin/shared/initialize"
	"douyin/shared/middleware"
	user "douyin/shared/rpc/kitex_gen/user/userservice"
	"log"

	"douyin/cmd/user/pkg/dal/mysql"
	"douyin/cmd/user/pkg/manager"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("user.srv")
	svr := user.NewServer(&UserServiceImpl {
			Manager: manager.Manager{
				Db: mysql.NewManager()},
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
