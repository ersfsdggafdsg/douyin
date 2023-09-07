package service

import (
	"context"
	"douyin/cmd/relation/pkg/dal/redis"
	"douyin/shared/rpc/kitex_gen/relation"
	"douyin/shared/utils/errno"
)

var Nil = redis.Nil
// 返回值：
// (true/false, error) 表示出现错误，为redis.Nil时需要写入缓存
// (true/false, nil)   表示查询到缓存的值
func queryCache(rdb *redis.RedisManager, ctx context.Context, req *relation.DouyinRelationActionRequest, resp errno.BaseResp) (result bool, err error) {
	// 设计这个返回值，有些犯难，查询的情况有下面这些
	// 1. 需要表示已经命中
	// 2. 需要表示没有命中
	// 3. 需要表示执行出错
	// 4. 命中了且没有关注
	// 5. 命中了且关注了
	// 但是golang中，返回nil后，一般表示可以继续往下执行的，
	// 这么做是无奈之举

	isFollowing, err := rdb.IsFollowing(req.UserId, req.ToUserId)
	switch err {
	case redis.Nil:
		// 需要读入到内存，无错误
		return false, redis.Nil
	case nil:
		return isFollowing, nil
	default:
		// 出错，返回错误
		errno.BuildBaseResp(errno.ServiceErrCode, resp)
		return false, err
	}
}
