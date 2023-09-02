package errhandler

import (
	"douyin/shared/rpc/kitex_gen/common"
	"douyin/shared/utils/errno"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)


func ErrorResponse(message string, err errno.ErrCode, c *app.RequestContext) {
	c.JSON(consts.StatusOK, &common.BaseResponse {
		StatusCode: int32(err),
		StatusMsg : message,
	})
}

func RPCCallErrorResponse(serviceName string, err errno.ErrCode, c *app.RequestContext) {
	ErrorResponse(fmt.Sprintf("'%s' not avaliable", serviceName), err, c)
}

func ParseTokenErrorResponse(err errno.ErrCode, c *app.RequestContext) {
	ErrorResponse("parse token error:", err, c)
}

func GenerateTokenErrorResponse(err errno.ErrCode, c *app.RequestContext) {
	ErrorResponse("generate token error:", err, c)
}
