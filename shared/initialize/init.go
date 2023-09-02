package initialize

import (
	"fmt"
	"log"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
)

func init() {
	fmt.Println("initialize: Setting logger level")
	klog.SetLevel(klog.LevelDebug)
	hlog.SetLevel(hlog.LevelDebug)
	log.SetFlags(log.Lmicroseconds | log.Llongfile)
}
