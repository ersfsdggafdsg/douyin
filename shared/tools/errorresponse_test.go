package tools

import (
	"douyin/shared/rpc/kitex_gen/comment"
	"douyin/shared/rpc/kitex_gen/feed"
	"fmt"
	"testing"
)

type myError struct {}
func (myError)Error() string {
	return "nothing"
}
func TestBuildBaseResp(t *testing.T) {
	fmt.Println("errorresponse.go")
	far := new(feed.DouyinFeedResponse)
	err := myError{}
	BuildBaseResp(err, 12, far)
	if far.GetStatusCode() != 12 {
		t.Log("FeedResponse failed")
	}
	car := new(comment.DouyinCommentListResponse)
	BuildBaseResp(err, 12, car)
	if car.GetStatusCode() != 12 {
		t.Log("FeedResponse failed")
	}
}
