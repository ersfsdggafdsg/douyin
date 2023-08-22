package initialize
import (
	"net"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/utils"
	consul "github.com/kitex-contrib/registry-consul"
)

// InitRegistry to init consul
func InitRegistry(srvName string) (registry.Registry, *registry.Info) {
	if len(os.Args) < 2 {
		klog.Fatalf("Missing argument: Port")
	}
	r, err := consul.NewConsulRegister("127.0.0.1:8500")
	if err != nil {
		klog.Fatalf("new consul register failed: %s", err.Error())
	}

	// Using snowflake to generate service name.
	sf, err := snowflake.NewNode(2)
	if err != nil {
		klog.Fatalf("generate service name failed: %s", err.Error())
	}
	info := &registry.Info{
		ServiceName: srvName,
		Addr:        utils.NewNetAddr("tcp", net.JoinHostPort("127.0.0.1", os.Args[1])),
		Tags: map[string]string{
			"ID": sf.Generate().Base36(),
		},
	}
	klog.Info("Service started:", info)
	return r, info
}

