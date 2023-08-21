package main

import (
	"douyin/shared/initialize"
	publish "douyin/shared/rpc/kitex_gen/publish/publishservice"
	"log"
	"os"

	"net"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("publish.srv")
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
