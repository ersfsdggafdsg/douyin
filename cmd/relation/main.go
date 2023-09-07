package main

import (
	"douyin/cmd/relation/pkg/mq"
	"douyin/cmd/relation/pkg/dal/mysql"
	"douyin/cmd/relation/pkg/dal/redis"
	"douyin/cmd/relation/pkg/manager"
	"douyin/shared/initialize"
	"douyin/shared/middleware"
	relation "douyin/shared/rpc/kitex_gen/relation/relationservice"
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("relation.srv")
	impl := RelationServiceImpl{
		manager.Manager{
			Db: mysql.NewManager(),
			Rdb: redis.NewManager(),
			Mq: mq.NewManager(),
		},
	}
	svr := relation.NewServer(&impl,
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
