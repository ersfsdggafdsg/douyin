package main

import (
	"douyin/shared/initialize"
	publish "douyin/shared/rpc/kitex_gen/publish/publishservice"
	"log"

	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("publish.srv")
	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
