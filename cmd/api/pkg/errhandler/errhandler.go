package errhandler

import (
	"douyin/shared/rpc/kitex_gen/common"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)


func ErrorResponse(message string, err error, errno int32, c *app.RequestContext) {
	hlog.Error(message, err)
	c.JSON(int(errno), &common.BaseResponse {
		StatusCode: errno,
		StatusMsg : err.Error(),
	})
}

func RPCCallErrorResponse(serviceName string, err error, errno int32, c *app.RequestContext) {
	ErrorResponse(fmt.Sprintf("rpc server '%s' failed:", serviceName), err, errno, c)
}

func ParseTokenErrorResponse(err error, errno int32, c *app.RequestContext) {
	ErrorResponse("parse token error:", err, errno, c)
}

func GenerateTokenErrorResponse(err error, errno int32, c *app.RequestContext) {
	ErrorResponse("generate token error:", err, errno, c)
}
