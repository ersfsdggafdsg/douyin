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

## rabbitmq安装

```shell
# On arch linux
sudo pacman -S rabbitmq
sudo rabbitmq-plugins enable --offline rabbitmq_peer_discovery_consul
```

TODO: 
1. 在网关添加中间件，过滤掉不存在的用户id、不存在的评论id、不存在的视频id
 uid: /user, /publish/list, /favorite/action, /comment/action, /comment/list, /favorite/list, /relation/follow/list, /relation/follower/list, /relation/friend/list,
 vid: /favorite/action, /comment/action, /comment/list
 uid, touid: /message/chat, /message/action, /relation/action
 uid(可以不存在): /feed

2. 提升
  - 更改jwt的模式，因为需要检查用户、视频、评论是否存在，使用jwt还不如直接服务端存一个验证用的
  - 更改idl文件，参数统一为一个结构体，这么做的好处是，可以直接通过统一的key生成函数，去查验数据是否存在
