package config
import (
	"gorm.io/gorm"
	"douyin/shared/initialize/clients/comment"
	"douyin/shared/rpc/kitex_gen/comment/commentservice"
	"douyin/shared/initialize/clients/favorite"
	"douyin/shared/rpc/kitex_gen/favorite/favoriteservice"
	"douyin/shared/initialize/clients/feed"
	"douyin/shared/rpc/kitex_gen/feed/feedservice"
	"douyin/shared/initialize/clients/message"
	"douyin/shared/rpc/kitex_gen/message/messageservice"
	"douyin/shared/initialize/clients/publish"
	"douyin/shared/rpc/kitex_gen/publish/publishservice"
	"douyin/shared/initialize/clients/relation"
	"douyin/shared/rpc/kitex_gen/relation/relationservice"
	"douyin/shared/initialize/clients/user"
	"douyin/shared/rpc/kitex_gen/user/userservice"
)
var Clients = struct {
	Comment  commentservice.Client
	Favorite favoriteservice.Client
	Feed     feedservice.Client
	Message  messageservice.Client
	Publish  publishservice.Client
	Relation relationservice.Client
	User     userservice.Client
} {
	Comment : comment.InitClient(),
	Favorite: favorite.InitClient(),
	Feed    : feed.InitClient(),
	Message : message.InitClient(),
	Publish : publish.InitClient(),
	Relation: relation.InitClient(),
	User    : user.InitClient(),
}

var DB *gorm.DB
