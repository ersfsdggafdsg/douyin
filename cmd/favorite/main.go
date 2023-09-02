package main

import (
	"douyin/shared/middleware"
	"douyin/shared/initialize"
	favorite "douyin/shared/rpc/kitex_gen/favorite/favoriteservice"
	"log"

	"douyin/cmd/favorite/pkg/model"
	"douyin/cmd/favorite/pkg/mysql"
	"github.com/cloudwego/kitex/server"
	"github.com/cloudwego/kitex/pkg/utils"
	"net"
	"os"
)

func main() {
	r, info := initialize.InitRegistry("favorite.srv")
	svr := favorite.NewServer(&FavoriteServiceImpl {
			Db: mysql.NewManager(initialize.InitMysql(
				"douyin", "zhihao", "douyin", &model.Favorite{},
			)),
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
