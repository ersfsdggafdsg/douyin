package rpc2http
import (
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/rpc/kitex_gen/base"
	"douyin/shared/utils"
)

func Video(v *rpc.VideoInfo) (*base.Video) {
	return &base.Video {
		Id           : v.Id,
		Author       : nil,
		PlayUrl      : utils.GetUrlFromHash(v.PlayHash),
		CoverUrl     : utils.GetUrlFromHash(v.CoverHash),
		FavoriteCount: v.FavoriteCount,
		CommentCount : v.CommentCount,
		Title        : v.Title,
	}
}

func User(u *rpc.UserInfo) (*base.User) {
	return &base.User {
		Id             : u.Id,
		Name           : u.Name,
		FollowCount    : u.FollowCount,
		FollowerCount  : u.FollowerCount,
		IsFollow       : u.IsFollow,
		Avatar         : u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature      : u.Signature,
		TotalFavorited : u.TotalFavorited,
		WorkCount      : u.WorkCount,
		FavoriteCount  : u.FavoriteCount,
	}
}
