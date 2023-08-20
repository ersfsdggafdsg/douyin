package main

import (
	relation "douyin/shared/rpc/kitex_gen/relation/relationservice"
	"log"
	"douyin/shared/initialize"
	"github.com/cloudwego/kitex/server"

)

func main() {
	r, info := initialize.InitRegistry("relation.srv")
	svr := relation.NewServer(
		new(RelationServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
