# tcp与udp区别
tcp是可靠的(流量控制与拥塞控制),面向连接的(点对点的),消耗资源多；

udp是不可靠的,无连接(点对点,广播),消耗资源少
# tcp为什么是三次握手,而不是两次
```
1试想两次握手,client发SYN报文给server,这个报文丢了；client重发了报文给server,建立连接,搞完业务需求,连接销毁；此时丢了的报文又找到了server,建立了连接,但是client认为连接不存在,server建立了一个假连接,浪费资源

2 通信的双方都必须维护一个序列号，以标识发出去的数据包中，哪些是已经被对方收到，三次握手的过程即是通信双方相互告知序列号起始值，并确认对方已经收到了序列号起始值的必经步骤，如果只是两次握手，至多只有连接发起方的起始序列号能被确认，另一方的序列号能不能被确认
```

# 为什么连接的时候是三次握手，关闭的时候却是四次挥手
```
连接的过程是
1 client发syn报文
2 server发ack报文表示client的syn报文收到了
3 server发syn报文
4 client发ack报文表示server的syn报文收到了
其中23一起,所以连接过程是三次握手

关闭的过程
1 client发fin报文
2 server发ack报文表示client的fin报文收到了
3 server发fin报文
4 client发ack报文表示server的fin报文收到了
关闭的时候23不能一起发,因为sever可能还有一些数据没有发给client
```
# 为什么最后要等一个TIME_WAIT时间呢？ TIME_WAIT状态要维持2MSL？
```
通常是client主动关闭
关闭过程
client established--(client没有数据要发送了发fin)-->fin_wait1--(收到server的ack)-->fin_wait2--(收到server的fin,发ack)-->time_wait-->close

server established--(收到client的fin,发ack)-->close_wait--(server没有数据要发送了发fin)-->last_ack-(收到client的ack)->close

1 确保client的报文使server端关闭,也就是说如果最后一个ack丢了还有时间重发
2 2MSL可以使得tcp两端的报文在网络中彻底消失,保证下一个tcp连接不会被本次连接干扰
```
# TCP在握手阶段怎么管理客户端的连接？
```
服务端维护了两个队列
半连接队列
client发syn到server,然后server发ack+syn,此时该连接就被放到半连接队列
全连接队列
第三次握手成功后,server将该连接放到全连接队列

全连接队列满的情况,看环境变量 tcp_abort_on_overflow,发rst给客户端还是忽略
SYN Flood 攻击时会造成服务端的半连接队列被占满，从而影响到服务
```

# TCP 通过哪些方式来保证数据的可靠性？
```
数据包:校验和
传输数据包:序列号,确认应答,超时重传
流量控制:拥塞控制

每个数据包都会有一个序列号,client发包到server会触发,server对这个包的确认,反过来也一样；那我们的传输模型就是一送一答?
不是的,发送方可以发多个包,一并等待确认,这就是发送窗口；发送窗口是min(接受方的接受窗口,拥塞窗口)

拥塞控制是 慢开始(指数增长)-{碰到慢开始门限}->拥塞避免(加法增长)-{网络拥塞}-{慢开始门限设置当前一半,拥塞窗口也是慢开始门限}->拥塞避免
```

# TCP长连接和短连接有什么区别？
```
短连接指的是双方一次读写就关闭,长连接指的是多次读写不关闭连接
长连接的优点是可以多次读写,缺点是对服务器压力大
短连接优点是管理简单,缺点是建立连接,销毁连接消耗资源
```
# TCP 粘包、拆包及解决方法？
```
tcp是面向字节流的
粘包是指要发送的数据包比较小跟其他数据包混在一起,无法区分
拆包是指发送的数据包比较大,需要拆分多次发送

解决方法: 在数据包头加上字节流的长度字段,或者约定某个特殊字符作为结尾,就知道数据的边界了
```


# 讲一下http
```
一段使用tcp协议传输的文本

浏览器请求=请求行+header+请求体
服务器响应=状态行+header+请求体

使用 curl -v {url} 可以看到请求相应的全过程

```
# HTTP 1.0和HTTP 1.1,HTTP2.0的主要区别是什么
```
1.0使用的是短连接,1.1默认使用长连接(connection:keep-alive)

1.1的请求响应模型是
请求1->请求2->请求3->响应1->响应2->响应3
响应必须按照请求的顺序来,哪怕响应1很慢,这会引起队头阻塞

2.0
支持多路复用的请求响应模型,不必按照顺序响应请求,不会有队头阻塞
可以推送信息到浏览器
```

# HTTPS
```
https = http+tls 
https使用混合加密,传输的内容使用的是对称加密,但是对称加密的密钥是服务器方的证书非对称加密来的

证书?

```

# HTTPS与HTTP区别
```
端口 80 vs 443
加密vs不加密

```

# http协议状态码
```
200 表示ok
4XX 客户端错
5XX 服务端错
```

# http常用字段
``` 
Access-Control-Allow-Origin 跨域相关
Content-Length
Connection:keep-alive 表示长链接
content-type 服务器回应时告诉客户端数据格式
```
# get与post区别
```
get没有请求body,想附带信息得放到url上面 querystring,例子 百度搜索的时候看url
get是幂等的,多次请求不会影响数据；post语义上不是,但具体看实现
get是读,可以做缓存
```