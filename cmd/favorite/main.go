package main

import (
	"douyin/shared/initialize"
	"douyin/shared/middleware"
	favorite "douyin/shared/rpc/kitex_gen/favorite/favoriteservice"
	"log"

	"douyin/cmd/favorite/pkg/dal/mysql"
	"douyin/cmd/favorite/pkg/dal/redis"
	"douyin/cmd/favorite/pkg/manager"
	"douyin/cmd/favorite/pkg/mq"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("favorite.srv")
	impl := FavoriteServiceImpl {
		Manager: manager.Manager{
			Db: mysql.NewManager(),
			Rdb: redis.NewManager(),
			Mq: mq.NewManager(),
		},
	}
	svr := favorite.NewServer(&impl,
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
		server.WithRegistry(r),
		server.WithMiddleware(middleware.ShowCallingMiddleware),
		server.WithRegistryInfo(info))

	impl.Mq.RunConsumers()
	
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
