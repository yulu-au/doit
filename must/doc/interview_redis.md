# redis-cluster与哨兵区别
```
哨兵解决的是高可用的问题,哨兵模式=主redis+从redis+哨兵防止redis挂
cluster解决高可用+水平扩展,三主三从+gossip协议

哨兵的缺点:不支持水平扩展,主redis只有一个,受限于内存,qps最多8w+
cluster的缺点:引入槽位的概念,keys */scan这种动作比较麻烦,原子操作也比较麻烦
```
# redis-cluster怎么获取所有的key
```
"keys *"在线上肯定是不行的,会引起明显的阻塞
cluster模式下,从节点默认是不提供读写服务的,可以自定义参考readonly命令
使用scan去遍历key,并且由于是cluster模式,需要在每个master上scan,然后汇总
```
## go-redis
```
ForEachMaster函数用来在cluster模式下执行需要对多个master执行的动作,比如获取所有key
```
## 从节点不提供读写
```
为什么不提供从节点的读,这是分担主库压力的标准动作,从cluster模式看,如果主库hold不住压力,最好的方式是加主redis的数量,也就是水平扩展
```

# redis持久化
```
生产环境方案
redis-cluster: master关闭持久化,slave打开RDB+AOF
```
## 数据恢复
```
将rdb文件放到数据文件位置就行

aof恢复,也是一样的
如果文件损坏 
redis-check-aof --fix {受损文件}
redis-check-dump {受损文件} 
```
## RDB redis database
```
为什么
复制内存这种方式较aof快,产生的文件较aof小,用来数据恢复较aof快

是什么
复制内存快照保存到硬盘上

怎么做
save(阻塞主线程)不推荐；bgsave(后台执行),redis就是这么做的,fork一个子进程去复制内存
配置文件里这么写
save 900 1 900秒内有一次动作就bgsave
save 300 10 同上
save 60 10000 同上


```
### 快照时发生数据修改怎么办
写时复制技术（Copy-On-Write, COW）,fork的子进程会共享父进程的内存空间,一旦父进程的内存发生改变,父进程就会将这段内存复制一份出来,子进程就看不到最新的改变

## AOF append only file
```
为什么
数据一致性,rdb方式可能会丢失数据

是什么
通过保存数据库执行的命令来记录数据库的状态

怎么做
打开AOF+配置写硬盘策略
配置文件 
appendonly yes
appendfsync everysec
auto-aof-rewrite-percentage 100
auto-aof-rewrite-min-size 64mb

写入策略取决于appendfsync参数
always表示每次写都写到硬盘
everysec 表示每秒刷一次 
no 表示由os决定
```
### aof重写
```
aof是记录命令这种日志形式的,所以文件会很大

重写主要是将多变一
set k1 v1
set k1 v2 -----> set k1 v3
set k1 v3
```

# 搭建redis-cluster
```
集群初始化命令
redis-cli --cluster create 192.168.10.10:6371 192.168.10.10:6372 192.168.10.10:6373 192.168.10.11:6374 192.168.10.11:6375 192.168.10.11:6376 --cluster-replicas 1

配置文件举例
port 6000
cluster-enabled yes
protected-mode no
cluster-config-file nodes.conf
cluster-node-timeout 5000
#对外ip
cluster-announce-ip 192.168.1.102 
cluster-announce-port 6000
cluster-announce-bus-port 16000
appendonly yes
```
# redis数据类型的底层实现
```
参考 https://github.com/zpoint/Redis-Internals/blob/5.0/README_CN.md

编码方式多样的原因是节省内存

string
底层有三种编码方式int或者embstr或者raw,就是说能转换为long类型的就使用int编码,除此以外占用字节数量少(44byte)的就使用embstr；此外使用sds编码方式的还有小类的SDS_TYPE_5,SDS_TYPE_8,SDS_TYPE_16,SDS_TYPE_32,SDS_TYPE_64



hash
底层数据结构是ziplist或者hashtable

list
底层实现是quicklist=ziplist+linkedlist

set
底层实现是intset或者hashtable；如果集合元素都是数值类型,并且元素数量少(512),使用intset

zset
底层实现是哈希表+ziplist或者skiplist；如果元素小且少使用ziplist
哈希表存score->member
```
# redis数据类型的使用场景
```
String(字符串)
二进制安全
可以包含任何数据,比如jpg图片或者序列化的对象,一个键最大能存储512M
---

Hash(字典)
键值对集合,即编程语言中的Map类型
适合存储对象,并且可以像数据库中update一个属性一样只修改某一项属性值(Memcached中需要取出整个字符串反序列化成对象修改完再序列化存回去)
存储、读取、修改用户属性


List(列表)
链表(双向链表)
增删快,提供了操作某一段元素的API
1、最新消息排行等功能(比如朋友圈的时间线) 2、消息队列


Set(集合)
哈希表实现,元素不重复
1、添加、删除、查找的复杂度都是O(1)  2、为集合提供了求交集、并集、差集等操作
1、共同好友 2、利用唯一性,统计访问网站的所有独立ip 3、好友推荐时,根据tag求交集,大于某个阈值就可以推荐


Sorted Set(有序集合)
将Set中的元素增加一个权重参数score,元素按score有序排列
数据插入集合时,已经进行天然排序
1、排行榜 2、带权重的消息队列

```

# 大KEY问题
```
redis的线上事故，原因是有个脚本删除了一个 redis 的大 key ，这个 key 是一个 zset 数据结构，里面有 1000w+ 数据，导致 cpu 100%.
如果一个大 key，del 会导致 cpu 飙升，那么给它一个过期时间，过期的那一刻也是产生同样的效果，等同于 del.

为什么cpu会满
元素数量多，实现上是 map+skiplist ，因为非数组结构（非连续内存），所以没法像操作单个元素那样删除所有元素，而是需要遍历删除每个节点，元素多，一个 op 整体删除肯定要阻塞其他请求较长时间。

解决方案
1 分批删除
比如 zremrangebyrank 一次删除 N 个，多次之间间隔 sleep 下（或者单次 op RTT 本来就有网络往返时间，一般不 sleep 也可以，看你们主业务的需要），因为是要删除的数据，删除慢点应该也无所谓，N 和是否需要 sleep 自己把握就行
2 unlink异步删除
4.0+ unlink 也可以，但是 unlink redis 线程之间通知之类的会多消耗一点，业务量大对 redis 的请求很频繁，用业务服务分批删、能替 redis 节省点性能可能对整个集群更划算，根据你们实际业务来判断
```

# 场景题
## zset怎么保证固定大小
```
zcard {key}
ZREMRANGEBYRANK {key} {start} {end} start end可以使用 -1 -2 指代倒数第几个元素
zadd {key} {score1} {member1} ...
```
## redis穿透问题
```
有一个表保存了点对点的关系。但是线上大多数点和点之间是没有关系的。所以 90%会打到数据库（如果 redis 不做缓存穿透的处理).
为了做缓存穿透的处理，在数据库查到为空时，在 redis 放了一个 key 等于参数，value 为空的值。但是，这部分值太多了。
解释一下大概就是人和人的点对点关系。可以理解为微博的好友添加,数据库里存了好友数据，业务发来两个用户查询两个人是否好友，大概率不是好友.

解决方案参考
1 redis 接入 mysql binlog 或 canal，让 redis 做全量 mysql 数据的同步
2 布隆过滤器定时重建

```
# 缓存DB一致性方案
```
1 cache aside方案(并发条件下会有脏数据)
更新数据库随后删除redis缓存内容

2延时双删(等待时间难以确认；脏数据还是有,因为DB从库比主库慢,读从库还是可能脏数据)
删除缓存,随后更新数据库,随后等待一小会再次删除缓存

3异步监听binlog删除+重试
是什么
抓取mysql的binlog得知更新,对应删除redis里的内容(删除失败走MQ重试)
优点
比较完善的方案,很多公司用
缺点
会有点慢,因为DB->binlog->分析->删除缓存
数据库拆分的时候,灰度切读流量的时候会有脏数据
```
# 缓存雪崩/击穿/穿透
```
经典redis三连
缓存击穿:请求到redis发现redis没有数据,于是请求到DB.(击穿是指击穿redis)
缓存穿透:请求到redis发现redis没有数据,于是请求到DB,也没有数据.(穿透是指DB也被击穿,穿透不应该存在,因为DB也没有的东西很明显是非法请求)
缓存雪崩:大量的缓存穿透

击穿经常发生,请求的数据redis没有去DB里拿是正常流程
穿透不应该发生,非法请求应该尽量避免
雪崩不应该发生,大量击穿会打挂DB

穿透的常见场景是非法请求,或者业务逻辑有问题.一般解决方案是布隆过滤器或者缓存空值
雪崩的常见场景是大量key在同一时间过期.一般解决方案是过期时间随机+-5%,保证在同一时刻不会有大量key过期
```

# 布隆过滤器
```
它是 一个二进制数组+多个hash函数 ,使用起来就是将想要缓存的值(DB里面的)经过hash运算加入到布隆过滤器里.
当请求携带"val"到redis里的时候,"val"也hash一下,布隆过滤器就知道它是不是在DB里了.

布隆过滤器有误差
当它判断"val"不在DB里,那一定不在；反之,判断"val"在DB里,有可能不在.
为什么有误差
"val"hash到某几个下标位置,这些位置都是1,那么bloomfilter判断"val"在DB里,但是这些"1"有可能是别的值hash到的

布隆过滤器的缺点
不能删除值,需要定时重建才能保证功能

```

# 面试题参考来源
https://easyhappy.github.io/travel-coding/mysql/%E5%89%8D%E8%A8%80.html
https://manbucoding.com/travel-coding/redis/%E5%89%8D%E8%A8%80.html