package main

import (
	"douyin/shared/middleware"
	"douyin/cmd/message/pkg/model"
	"douyin/cmd/message/pkg/mysql"
	"douyin/shared/initialize"
	message "douyin/shared/rpc/kitex_gen/message/messageservice"
	"log"
	"os"
	"net"
	"github.com/cloudwego/kitex/pkg/utils"

	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("message.srv")
	svr := message.NewServer(&MessageServiceImpl{
		Db: mysql.NewManager(initialize.InitMysql(&model.Message{}))},
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
