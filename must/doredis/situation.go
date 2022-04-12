package doredis

/*
使用场景
https://cloud.tencent.com/developer/article/1867518
*/

/*
排行榜相关
app里有很多课程,这些课程被观看的时长是不一样的,需要实时给出被观看时长最大的十个课程
*/

/*
全局ID
int类型，incrby，利用原子性

incrby userid 1000

分库分表的场景，一次性拿一段
*/

/*
计数器

int类型，incr方法

例如：文章的阅读量、微博点赞数、允许一定的延迟，先写入Redis再定时同步到数据库
*/

/*
int类型，incr方法

以访问者的ip和其他信息作为key，访问一次增加一次计数，超过次数则返回false
*/

/*
使用 Redis 统计在线用户人数
*/

/*
用户消息时间线timeline

list，双向链表，直接作为timeline就好了。插入有序
*/

/*
点赞、签到、打卡
*/
