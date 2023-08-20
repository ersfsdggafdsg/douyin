package main

import (
	comment "douyin/shared/rpc/kitex_gen/comment/commentservice"
	"log"
	"douyin/shared/initialize"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("comment.srv")
	svr := comment.NewServer(
		new(CommentServiceImpl),
		server.WithRegistry(r),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
