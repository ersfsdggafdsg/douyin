package initialize

import (
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

// InitRegistry to init consul
func InitRegistry() (registry.Registry, *registry.Info) {
	// build a consul client
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		hlog.Fatalf("new consul client failed: %s", err.Error())
	}

	r := consul.NewConsulRegister(consulClient)

	if err != nil {
		hlog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: "api.srv",
		Addr: utils.NewNetAddr("tcp", "127.0.0.1:8888"),
		Weight: registry.DefaultWeight,
		Tags: nil,
	}
	return r, info
}

