package middleware
import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ShowCallingMiddleware

type args interface {
	GetFirstArgument() interface{}
}

type result interface {
	GetResult() interface{}
}

func ShowCallingMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		if ri.To().Method() != "PublishAction" {
			// 避免输出的内容过多。一个视频文件会很大
			klog.Infof("real request: %+v\n", req.(args).GetFirstArgument())
		}
		// get local service information
		klog.Infof("local service name: %v\n", ri.From().ServiceName())
		// get remote service information
		klog.Infof("remote service name: %v, remote method: %v\n", ri.To().ServiceName(), ri.To().Method())
		if err := next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v\n", resp.(result).GetResult())
		return nil
	}
}

