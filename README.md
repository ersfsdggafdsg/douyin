# douyin

## 目录结构
```
 cmd
 +- api API网关，使用Hertz
 |  +- biz Hertz自动生成的路径，需要实现其中的handler
 |  +- script router_gen.go router.go main.go build.sh 由Hertz根据protobuf生成
 |  +- output api 由build.sh编译产生
 |  +- api-gen.sh 自动编译protobuf的脚本，由Afeather2017编写
 |  +- init 初始化
 |  |  `- rpc 初始化rpc服务
 |  `- pkg handler的具体实现
 +- comment
 |  +- pkg 
 |  |  `- mysql 数据库操作的实现
 |  +- init
 |  `- config 初始化使用的代码
 |     +- rpc rpc实现
 |     +- ...init.go rpc初始化
 |     `- init.go 初始化
 +- favorite
 +- feed
 +- message
 +- publish
 +- relation
 +- storage
 `- user
```

## cmd/api

如果更新了`idl/http`下的protobuf文件，一定要执行`api-gen`来更新API网关
