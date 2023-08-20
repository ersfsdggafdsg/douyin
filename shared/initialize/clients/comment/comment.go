package comment

import (
	"douyin/shared/rpc/kitex_gen/comment/commentservice"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	consul "github.com/kitex-contrib/registry-consul"
)

func InitClient() (commentservice.Client) {
	// init resolver
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatalf("new consul client failed: %s", err.Error())
	}

	// create a new client
	c, err := commentservice.NewClient(
		"comment.srv",
		client.WithResolver(r),                                     // service discovery
		client.WithLoadBalancer(loadbalance.NewWeightedBalancer()), // load balance
		client.WithMuxConnection(1),                                // multiplexing
	)
	if err != nil {
		klog.Fatalf("ERROR: cannot init client: %v\n", err)
	}
	return c
}



