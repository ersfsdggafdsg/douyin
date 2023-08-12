// Code generated by hertz generator. DO NOT EDIT.

package publish

import (
	publish "douyin/cmd/api/biz/handler/publish"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			_publish.POST("/action", append(_publishactionMw(), publish.PublishAction)...)
			_publish.GET("/list", append(_publishlistMw(), publish.PublishList)...)
		}
	}
}
