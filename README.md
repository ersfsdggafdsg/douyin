# douyin

## rabbitmq安装

```shell
# On arch linux
sudo pacman -S rabbitmq
sudo rabbitmq-plugins enable --offline rabbitmq_peer_discovery_consul
```

## consul安装
``` shell
# On arch linux
sudo pacman -S consul
```

## 编译并执行
``` shell
make run
```

## 编译的时候可能会遇到的问题

1. 编译的时候报错说找不到AVCodecContext。这需要你修改thumbnailer的代码，详情见shared/utils/cover/cover.go。其实我很想自己写一个的，但是没时间了。
2. 在1024code上可以编译但是没有办法获取到视频。这需要你取消掉cmd/api/router.go的注释。详情也写在里面，总之就是要反向代理。
3. Mysql、redis、rabbitmq、consul等无法链接。这需要你修改run.sh，改掉环境变量。
4. 视频无法保存。这需要你创建cmd/storage/static文件夹。我可能已经创建了。

## 项目的目录结构

```
 项目根目录
 ├╴idl           idl接口描述文件
 ├╴tools         一些供调试的工具
 ├╴makefile      全局makefile
 ├╴cmd           各个微服务和网关
 ╰╴shared        复用的代码
   ├╴middleware  通用的中间件
   ├╴rpc         kitex的生成内容
   ├╴config      配置RPC调用客户端
   ├╴initialize  初始化，这个是供微服务使用的
   ├╴consts      常量
   ╰╴utils       一些通用的组件
```

```
 favorite      cmd下的各个微服务 
 ├╴handler.go  将功能实现转向service
 ├╴pkg  
 │ ├╴service   handler的实现，有每个路由的实现以及RPC调用的实现
 │ │ ├╴favoriteadd.go 
 │ │ ╰╴...
 │ ├╴mq        消息队列实现
 │ │ ╰╴mq.go 
 │ ├╴dal       数据访问层
 │ │ ├╴mysql  
 │ │ │ ╰╴mysql.go 
 │ │ ├╴model  
 │ │ │ ╰╴model.go 
 │ │ ╰╴redis  
 │ │   ╰╴redis.go 
 │ ╰╴manager   管理器，管理了消息队列和数据访问层
 │   ╰╴manager.go 
 ├╴main.go 
 ╰╴makefile
```

```
 cmd/api 
 ├╴api-gen.sh 
 ├╴router_gen.go 
 ├╴pkg           API网关专用的一些组件，包括JWT、中间件
 ├╴initialize    初始化，这个供api网关使用的
 ├╴main.go 
 ├╴router.go     路由。由于要部署到1024code上，所以有一个反向代理
 ├╴biz  
 │ ├╴handler     各个路由的实现
 │ ├╴router      各个路由的一些配置，比如中间件
 │ ╰╴model       各个路由使用的结构体
 ╰╴makefile
```

## 项目说明

见https://pa3l7zekcrz.feishu.cn/docx/NW5Pd8slsoif16xt0ZBcnA8snhg
