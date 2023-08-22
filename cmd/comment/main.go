package main

import (
	comment "douyin/shared/rpc/kitex_gen/comment/commentservice"
	"log"
	"os"
	"net"
	"github.com/cloudwego/kitex/pkg/utils"
	"douyin/shared/initialize"
	"github.com/cloudwego/kitex/server"
)

func main() {
	r, info := initialize.InitRegistry("comment.srv")
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithRegistry(r),
		server.WithServiceAddr(utils.NewNetAddr("tcp",
			net.JoinHostPort("127.0.0.1", os.Args[1]))),
		server.WithRegistryInfo(info))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
