package main

import (
	"douyin/shared/initialize"
	feed "douyin/shared/rpc/kitex_gen/feed/feedservice"
	"log"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("feed.srv")
	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
