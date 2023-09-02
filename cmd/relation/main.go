package main

import (
	"douyin/shared/middleware"
	"douyin/shared/initialize"
	relation "douyin/shared/rpc/kitex_gen/relation/relationservice"
	"log"
	"net"
	"os"

	"douyin/cmd/relation/pkg/model"
	"douyin/cmd/relation/pkg/mysql"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("relation.srv")
	svr := relation.NewServer(&RelationServiceImpl{
			Db: mysql.NewManager(initialize.InitMysql(
				"douyin", "zhihao", "douyin", &model.Relation{})),
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
