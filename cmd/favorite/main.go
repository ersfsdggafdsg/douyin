package main

import (
	favorite "douyin/shared/rpc/kitex_gen/favorite/favoriteservice"
	"log"
	"douyin/shared/initialize"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("favorite.srv")
	svr := favorite.NewServer(
		new(FavoriteServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
