# grpc服务端的启动流程
```
1 初始化tcp连接
2 初始化grpcserver
3 注册自定义service到grpcserver
4 grpcserver监听tcp连接处理请求
```

# grpc服务类型有哪些
```
简单rpc
服务端流
客户端流
双端流
```

# keep-alive是针对stream还是connection
```
每个grpc请求都是 stream,Keepalive能够让grpc的每个stream保持长连接状态,适合一些执行时间长的请求

grpc源代码下二级目录有keepalive包,EnforcementPolicy,ServerParameters是用来配置keepalive的,配置grpcserver的时候传入这些

https://github.com/grpc/grpc-go/blob/master/examples/features/keepalive/server/main.go

对于http协议也有这个东西,不过是另外的意思了
参考 
https://pandaychen.github.io/2020/09/01/GRPC-CLIENT-CONN-LASTING/
http://xingyys.tech/grpc2/
```

# http2 conn的keepalive处理流程

# grpc通信报文格式

# 常见的拦截器有哪些,第三方
```
go-grpc-middleware
```

# 多路复用指什么
```
http2的多路复用
```
# 传输报文中metadata存储什么内容
```
使用场景类型http的header
```

# 如何自定义resolver
```
实现builder和resolver接口,builder是用来生成解析器的,resolver是解析器
resolver接口更像是提示而非必要实现,以ETCD做服务发现为例,在Build方法里监听key,更新服务地址
```

# 如何自定义balance

# 如何实现grpc全链路跟踪
# 客户端connection连接状态有哪些
# 客户端如何拿到服务端的服务函数List
```
proto文件生成的stub代码
```
# grpc什么是backoff协议
```
https://grpc.github.io/grpc/core/md_doc_connection-backoff.html
当某个连接失败了,不要立刻重试,因为这可能加重网络阻塞,或者使得服务端更加挂,需要退避算法.
将真,这有点像tcp的拥塞控制
```
# grpc如何为每个stream进行限流,什么是 flow control
```

```
# 什么是HPack
```
http2使用的压缩算法,用来压缩http-header,节省消息头使用的流量
```