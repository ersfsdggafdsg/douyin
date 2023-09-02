package middleware

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Log(ctx context.Context, c *app.RequestContext) {
	hlog.Debug(ctx, " ", string(c.GetRequest().RequestURI()))
}
