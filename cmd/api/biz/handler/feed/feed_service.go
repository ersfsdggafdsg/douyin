// Code generated by hertz generator.

package feed

import (
	"context"

	feed "douyin/cmd/api/biz/model/feed"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Feed .
// @router /douyin/feed [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req feed.DouyinFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(feed.DouyinFeedResponse)

	c.JSON(consts.StatusOK, resp)
}
