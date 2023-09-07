package main

import (
	"douyin/cmd/comment/pkg/dal/mysql"
	"douyin/cmd/comment/pkg/manager"
	"douyin/cmd/comment/pkg/mq"
	"douyin/shared/initialize"
	"douyin/shared/middleware"
	comment "douyin/shared/rpc/kitex_gen/comment/commentservice"
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("comment.srv")
	impl := CommentServiceImpl {
		manager.Manager{
			Db: mysql.NewManager(),
			Mq: mq.NewManager(),
		},
	}
	svr := comment.NewServer(&impl,
		server.WithRegistry(r),
		server.WithMiddleware(middleware.ShowCallingMiddleware),
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
		server.WithRegistryInfo(info))

	impl.Mq.RunConsumers()

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
