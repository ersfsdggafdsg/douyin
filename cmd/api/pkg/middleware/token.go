package middleware

import (
	"douyin/cmd/api/pkg/errhandler"
	"douyin/shared/utils"
	"douyin/shared/utils/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/net/context"
)

// 尽可能确保能拿到token的函数
func getToken(c *app.RequestContext) string {
	// get请求可以使用c.Query
	// 但是post-form不行
	token := c.Query("token")
	if token == "" {
		token, _ = c.GetPostForm("token")
	}
	// hlog.Infof("token: '%s'", token)
	return token
}

func parseToken(token string) (int64, error) {
	uid := int64(-1)
	if token != "" {
		claims, err := utils.ParseToken(token)
		if err != nil {
			// Token有错误或已经过期
			return -1, err
		}
		uid = int64(claims.Id)
	}
	return uid, nil
}

// jwt检查的中间件，提供给不管有没有登录都可以用的服务
func ParseTokenAndContinue(ctx context.Context, c *app.RequestContext) {
	token := getToken(c)
	uid, err := parseToken(token)
	if err != nil {
		// 出错了就结束
		hlog.Error("parse token:", err)
		errhandler.ErrorResponse("Token is wrong or expired",
			errno.AuthorizationFailedErrCode, c)
		c.Abort()
		return
	}
	ctx = context.WithValue(ctx, "uid", uid)
	c.Next(ctx)
}

// jwt检查的中间件，给必须登录的服务使用
func NoOrWrongTokenAbort(ctx context.Context, c *app.RequestContext) {
	token := getToken(c)
	uid, err := parseToken(token)
	if err != nil {
		// 出错了就结束
		hlog.Error(err)
		errhandler.ErrorResponse("Token is wrong or expired",
			errno.AuthorizationFailedErrCode, c)
		c.Abort()
		return
	}
	if uid <= 0 { // 未登录用户
		hlog.Error(err)
		errhandler.ErrorResponse("Not login",
			errno.AuthorizationFailedErrCode, c)
		c.Abort()
		return
	}
	ctx = context.WithValue(ctx, "uid", uid)
	c.Next(ctx)
}

// jwt检查的中间件，给同时需要保证userid和token带的id相同的服务使用
func TokenUserIdSame(ctx context.Context, c *app.RequestContext) {
	token := getToken(c)
	uid, err := parseToken(token)
	if err != nil {
		// 出错了就结束
		hlog.Error(err)
		errhandler.ErrorResponse("Token is wrong or expired",
			errno.AuthorizationFailedErrCode, c)
		c.Abort()
		return
	}
	if uid <= 0 { // 未登录用户
		hlog.Error(err)
		errhandler.ErrorResponse("Not login",
			errno.AuthorizationFailedErrCode, c)
		c.Abort()
		return
	}
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		hlog.Error(err)
		errhandler.ErrorResponse("Wrong Parameter",
			errno.ParamErrCode, c)
		c.Abort()
		return
	}
	if uid != id {
		hlog.Error(err)
		errhandler.ErrorResponse("Wrong token",
			errno.AuthorizationFailedErrCode, c)
		c.Abort()
		return
	}
	ctx = context.WithValue(ctx, "uid", uid)
	c.Next(ctx)
}
