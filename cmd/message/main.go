package main

import (
	message "douyin/shared/rpc/kitex_gen/message/messageservice"
	"log"
	"douyin/shared/initialize"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("message.srv")
	svr := message.NewServer(
		new(MessageServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
