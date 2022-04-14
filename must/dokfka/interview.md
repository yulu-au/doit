# 生产者消息如何选择到哪个分区
通常业务上想要保证消息的有序性,有两种方案.1 一个topic只有一个分区 2 想要保证顺序的消息使用一个key+生产者的负载均衡策略不能是轮询或者随机
## kafka-go
1 生产者的balance策略选择能使得相同的key放到一个分区里的,比如hash;
2 生产者发消息一类业务指定一个key；
# 消费者组里放几个消费者

一般是跟分区数量一致,多出来的消费者无法消费分区;
因为一个消费者组只能对同一分区消费一次

# 对kafka的理解
## zookper
Kafka通过 ZooKeeper 管理集群配置、选举 Leader 以及在 consumer group 发生变化时进行 Rebalance；
目前kafka会把offset保存在kafka的"consumer_offset"topic
## topic
主题,可以看作一个标签,同一主题的消息一起处理
### partition
分区,主题的子概念,同一主题可以分为多个分区；分区其实就是分片,那么同一主题所有分区加起来就是数据的全集,一旦一个分区挂了整个数据都不可用,所以又要做冗余的副本,副本不提供读写服务,仅用来做后备；真猴子(leader分区) 后备猴子(follower分区)
#### segment
对分区的再一次切分,因为生产者不断写数据会导致分区文件的长度不断增大
#### offset
偏移量,消费者消费某个分区的时候得知道自己消费到哪里了,这个偏移量就是记录消费的位置的；
#### lag
这个用来表示消费者落后生产者多远；一般健康的系统它无限接近0；如果它很大,表示消费者跟不上生产者的进度,消息挤压的多,因为新的消息会在filecache里,旧的消息大概会到磁盘那里,这样消费的时候就会从磁盘里把消息拿出来,那么消费的速度就会越来越慢,lag会越来越大
## broker
kafka的节点,分为leader节点和follower节点
## 消费者组
消费者组里包含多个消费者；同一个消费者组里所有的消费者对同一分区记录一个offset,也就是说对某个分区而言,同一个消费者组里的消费者是一体的,只能读取一次这个分区；

实践中,消费者组中消费者的数量与对应主题分区的数量是相同的,这样保证每个分区都有一个对应的消费者；如果某topic有三个分区,对应的消费者组有四个消费者,那就会有一个消费者空转
# kafka为什么性能高
## 顺序读写
kafka将消息记录持久化到本地磁盘,它将message不断追加到文件尾,这样相比随机读写会快；
所以它不删除message,每个消费者对每个分区持有offset去消费；
### 日志删除策略
所以它的硬盘占用会很快满,删除数据的策略一是基于时间,二是基于partition大小,具体见配置文档； 
log.retention.bytes  log.retention.hours

### 读
kafka存的message会产生两类文件,索引文件,数据文件是分区+分段,而且不论是索引还是数据都是有序的,这样就给二分查找

## page cache
使用os自己的page cache,而不是jvm的堆内存免去gc的烦恼,os对page cache也会有很多优化
## 零拷贝
避免了在内核空间和用户空间之间的拷贝
## sendfile
系统调用,允许pagecache直接拷贝数据到socket缓冲区,而不需要经过用户缓冲区的转换
## 分区分段
读某个topic的某个分区的数据,对于一个分区来讲,它保存在不同的broker上,这样io的压力会分散开
## 批量读写,压缩
一次读大量的数据,节约资源；kafka生产者写入数据不是实时发给消费者,这就是原因

# 消息队列的两种模型
# 队列模型
一个队列
# 发布订阅模型
一个消息可以被多个消费者消费
## 消费者组
一个消费者组对同一个分区只能消费一次；如果所有的消费者都在同一个组里,那对于所有的消息而言,只能被一个消费者消费一次；
反之,如果每个消费者都创建一个自己的组,那它们的消费offset就是互不影响的了

# 一些可视化工具
offset exploer
# 配置文件地址
find / -name "server.properties"
# 数据文件所在
配置文件里的配置项
log.dirs=/kafka/kafka-logs-66495db08205；

文件的组织形式就是索引文件加上数据文件,当然这两类文件都被切分为多个,而且都是有序的,二分查找可以使用

# kafka通信模型
生产者push消息到broker,消费者pull消息到本地；
分区存在replication的时候会发生副本之间的拷贝

# kafka 分区同步机制
分区的N个replicas中。其中一个replica为 leader，其他都为follower；leader处理partition的所有读写请求，与此同时，follower会被动定期地去复制leader上的数据。
## ISR(In-Sync Replicas)
是一个集合,里面是要与leader分片保持一直的follower分片；follower分片必须能与zookeeper保持会话(心跳机制),follower分片本能复制leader上的所有写操作，并且不能落后太多,replica.lag.time.max.ms；所有的副本(replicas)统称为Assigned Replicas，即AR,ISR是AR中的一个子集
### 机制
leader新写入的消息，consumer不能立刻消费，leader会等待该消息被所有ISR中的replicas同步后更新HW，此时消息才能被consumer消费。这样就保证了如果leader所在的broker失效，该消息仍然可以从新选举的leader中获取
### 优点
Kafka的复制机制既不是完全的同步复制，也不是单纯的异步复制。事实上，同步复制要求所有能工作的follower都复制完，这条消息才会被commit，这种复制方式极大的影响了吞吐率。而异步复制方式下，follower异步的从leader复制数据，数据只要被leader写入log就被认为已经commit，这种情况下如果follower都还没有复制完，落后于leader时，突然leader宕机，则会丢失数据。而Kafka的这种使用ISR的方式则很好的均衡了确保数据不丢失以及吞吐率。

# kafka为什么要分区
分区其实就是数据分片机制,在不同的中间件里都有体现,比如redis的槽位,比如es的分片；
数据分片的好处是水平扩展方便,这样能hold住更大的吞吐量

# 消费者重平衡
## 怎么看消费者组里有几个消费者
其实是看分区和消费者的对应关系(CONSUMER-ID与PARTITION)
```
kafka-consumer-groups.sh  --bootstrap-server 172.20.0.1:9092 --describe --group g2

GROUP           TOPIC           PARTITION  CURRENT-OFFSET  LOG-END-OFFSET  LAG             CONSUMER-ID                                                                                  HOST            CLIENT-ID
g2              test            1          44              44              0               consume@yu-LENOVO-G480L (github.com/segmentio/kafka-go)-a2fdf392-673d-4d9e-84c0-5dcf1ddbbff3 /172.20.0.1     consume@yu-LENOVO-G480L (github.com/segmentio/kafka-go)
g2              test            2          34              34              0               consume@yu-LENOVO-G480L (github.com/segmentio/kafka-go)-a2fdf392-673d-4d9e-84c0-5dcf1ddbbff3 /172.20.0.1     consume@yu-LENOVO-G480L (github.com/segmentio/kafka-go)
g2              test            0          45              45              0               consume@yu-LENOVO-G480L (github.com/segmentio/kafka-go)-1167d694-de3f-43dd-84f6-8adb00d754e1 /172.20.0.1     consume@yu-LENOVO-G480L (github.com/segmentio/kafka-go)

```
分区1,2对应一个消费者,分区0对应另外一个消费者；这样很明显消费者组只有俩消费者,一个消费者多领了一个分区

如果分区1,2,3分别对应一个消费者,那有可能存在一个消费者没有得到分区
## 重平衡为什么不好
重平衡其实就是一个协议，它规定了如何让消费者组下的所有消费者来分配topic中的每一个分区。比如一个topic有100个分区，一个消费者组内有20个消费者，在协调者的控制下让组内每一个消费者分配到5个分区，这个分配的过程就是重平衡；
因为重平衡过程中，消费者无法从kafka消费消息，这对kafka的TPS影响极大，而如果kafka集内节点较多，比如数百个，那重平衡可能会耗时极多。数分钟到数小时都有可能，而这段时间kafka基本处于不可用状态。
## 触发条件
消费者与分区的比例关系被打破
```
1 消费者数量增减,有的消费者挂了
2 主题的分区数量更改
3 订阅的主题发生变化，当消费者组使用正则表达式订阅主题，而恰好又新建了对应的主题，就会触发重平衡
```
## 重平衡策略
### Sticky
听名字就知道，主要是为了让目前的分配尽可能保持不变，只挪动尽可能少的分区来实现重平衡
### RoundRobin
RoundRobin是基于全部主题的分区来进行分配的，同时这种分配也是kafka默认的rebalance分区策略
### Range
这种分配是基于每个主题的分区分配，如果主题的分区分区不能平均分配给组内每个消费者，那么对该主题，某些消费者会被分配到额外的分区

## 避免重平衡
首先要知道，如果消费者真正挂掉了，那我们是没有什么办法的，但实际中，会有一些情况，会让kafka错误地认为一个正常的消费者已经挂掉了，我们要的就是避免这样的情况出现；

三个参数，session.timout.ms控制心跳超时时间，heartbeat.interval.ms控制心跳发送频率，以及max.poll.interval.ms控制poll的间隔

# kafka保证可靠性
## 分区多副本架构
多个副本保证即使挂了一个也能提供数据服务
## 发送消息
ack
```
0 意味着如果生产者能够通过网络把消息发送出去，那么就认为消息已成功写入 Kafka

1 意味若 Leader 在收到消息并把它写入到分区数据文件（不一定同步到磁盘上）时会返回确认或错误响应；这个模式下仍然有可能丢失数据，比如消息已经成功写入 Leader，但在消息被复制到 follower 副本之前 Leader发生崩溃。

all -1
意味着 Leader 在返回确认或错误响应之前，会等待所有同步副本都收到悄息；如果和 min.insync.replicas 参数结合起来，就可以决定在返回确认前至少有多少个副本能够收到悄息，生产者会一直重试直到消息被成功提交
```
# kafka数据一致性
数据一致性主要是说不论是老的 Leader 还是新选举的 Leader，Consumer 都能读到一样的数据

```
图
replica0 replica1 replica2

Message1 Message1 Message1
Message2 Message2 Message2
-------------------------- High Water Mark
Message3 Message3
Message4 

```
假设分区的副本为3，其中副本0是 Leader，副本1和副本2是 follower，并且在 ISR 列表里面。虽然副本0已经写入了 Message4，但是 Consumer 只能读取到 Message2。因为所有的 ISR 都同步了 Message2，只有 High Water Mark 以上的消息才支持 Consumer 读取，而 High Water Mark 取决于 ISR 列表里面偏移量最小的分区，对应于上图的副本2，这个很类似于木桶原理

# 不重复消费
重复消费的可能场景:消费者消费一条消息还没等到commit offset就挂了,然后消费者重启再次消费
业务有幂等性,这样的话重复消费消息也不会有问题
# kafka消息有序性
分区内有序
```
一个topic只给一个分区
生产者生产消息放到分区的策略选择hash+消息带key
```

# 坑
## kafka-go
```
https://cloud.tencent.com/developer/article/1809467
消息丢失
func (r *Reader) ReadMessage(ctx context.Context) (Message, error)
假失败
```