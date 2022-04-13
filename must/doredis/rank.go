package doredis

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/go-redis/redis/v8"
)

/*
使用场景汇总
https://cloud.tencent.com/developer/article/1867518
*/

/*
需求
  排行榜相关
  app里有很多课程,这些课程{学习资源}被点赞的数量是不一样的,需要实时给出被点赞最大的十个课程

具体动作
  初始化
    最开始将所有课程ID加入到"rank-course"里,这是一个zset;新课程需要后台运营调用自己的接口将课程ID加入到"rank-course"
	"zadd rank-course 100000 course-1001 ..."
  更新member对应的score,score代表点赞数据
    "ZINCRBY rank-course -100 course-1001"
  获取top100的课程ID
    "ZREVRANGE rank-course 0 99 withscores"
  获取课程信息全集
    通常zset里存ID,其他信息放在另外的结构,比如hashmap,或者直接DB

持久化
  数据异步刷到数据库,定时检查二者
特别注意
  分数相同的问题,score可能是一样的;这种情况排序是按照member的字典序来的,如果想要按照member加入的时间来区分这种score一样的情况
  那就{新socre}={原本的score}<<11+时间戳,但是要注意score字段精确表达的整数范围为-2^53到2^53;或者还可以用时间戳魔改member字段



点赞数量相同的话按照元素的字典序展示
*/
func InitRankCourse() {
	//init
	members := make([]*redis.Z, 0)
	for i := 0; i < 100; i++ {
		//course-{x}表示课程id,score表示观看时长
		z := redis.Z{Score: 1000 * rand.Float64(), Member: fmt.Sprintf("course-%v", i)}
		members = append(members, &z)
	}

	addZset("rank-course", members)
}

//课程的时长更新
func AddScore(key string, incr float64, member string) {
	_, err := ConnCluster.ZIncrBy(context.Background(), key, incr, member).Result()
	if err != nil {
		log.Panicln(err)
	}
}

//取出时间最长的课程ID
func Top10(key string) []redis.Z {
	res, err := ConnCluster.ZRevRangeWithScores(context.Background(), key, 0, 9).Result()
	if err != nil {
		log.Panicln(err)
	}

	return res
}

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
点赞、签到、打卡
*/
