package main

import (
	"douyin/shared/middleware"
	"douyin/cmd/comment/pkg/model"
	"douyin/cmd/comment/pkg/mysql"
	"douyin/shared/initialize"
	comment "douyin/shared/rpc/kitex_gen/comment/commentservice"
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("comment.srv")
	svr := comment.NewServer(&CommentServiceImpl {
			Db: mysql.NewManager(initialize.InitMysql(
				"douyin", "zhihao", "douyin", &model.Comment{},
			)),
		},
		server.WithRegistry(r),
		server.WithMiddleware(middleware.ShowCallingMiddleware),
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
