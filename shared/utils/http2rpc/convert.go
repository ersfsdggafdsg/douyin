package http2rpc
import (
	"douyin/shared/rpc/kitex_gen/rpc"
	"douyin/shared/rpc/kitex_gen/base"
)

func Video(v *base.Video) (*rpc.VideoInfo) {
	return &rpc.VideoInfo {
		VideoId      : v.Id,
		AuthorId     : v.Author.Id,
		PlayUrl      : v.PlayUrl,
		CoverUrl     : v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount : v.CommentCount,
		Title        : v.Title,
	}
}

func User(u *base.User) (*rpc.UserInfo) {
	return &rpc.UserInfo {
		UserId         : u.Id,
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
