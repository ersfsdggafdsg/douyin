package main

import (
	"douyin/shared/middleware"
	"douyin/shared/initialize"
	feed "douyin/shared/rpc/kitex_gen/feed/feedservice"
	"log"
	"os"
	"net"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("feed.srv")
	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithRegistry(r),
		server.WithMiddleware(middleware.ShowCallingMiddleware),
		server.WithRegistryInfo(info),
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
		)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
