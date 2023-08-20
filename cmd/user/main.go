package main

import (
	user "douyin/shared/rpc/kitex_gen/user/userservice"
	"log"
	"douyin/shared/initialize"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("user.srv")
	svr := user.NewServer(
		new(UserServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
