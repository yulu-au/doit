//收集好看的博客
//https://www.lixueduan.com/categories/Redis/
package doredis

import (
	"context"
	"fmt"
	"log"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	ConnCluster = newRedisConn()
)

func newRedisConn() *redis.ClusterClient {
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	//这里至填写了master的位置,因为redis-cluster中slave节点默认不提供读写服务
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"127.0.0.1:6000", "127.0.0.1:6001", "127.0.0.1:6002"},
	})
	return rdb
}

//keys 与 scan
/*
key很多的情况下
keys pattern 会阻塞redis
scan因为有游标的功能不用一次性返回所有,于是成为这种情况的解决方案.
另外count不能太大,不然很多也会卡住redis.不能太小,多次操作使得网络延迟也变多
*/
//连接单个redis的情况
func PrintAllKeysFromSingle() {
	all := make([]string, 0)
	var cur uint64 = 0
	for {
		time.Sleep(time.Millisecond)
		res, cursor, err := ConnCluster.Scan(context.Background(), cur, "*", 5).Result()
		if err != nil {
			panic(err)
		}
		all = append(all, res...)
		if cursor == 0 {
			break
		}
		cur = cursor
	}
	for _, v := range all {
		fmt.Println(v)
	}
}

//连接集群的情况
func PrintAllKeysFromCluster() []string {
	res := make([]string, 0)

	keys := func(ctx context.Context, c *redis.Client) error {
		var cursor uint64
		ks, cursor, err := c.Scan(context.Background(), cursor, "*", 100).Result()
		if err != nil {
			return err
		}
		res = append(res, ks...)
		for cursor != 0 {
			ks, cursor, err = c.Scan(context.Background(), cursor, "*", 100).Result()
			res = append(res, ks...)
		}
		return nil
	}
	ConnCluster.ForEachMaster(context.Background(), keys)

	return res
}

//这里其实要考虑不同类型value,如果是hashmap类型的val,使用get去哪会报错
//api是对redis命令的简单封装,那么不同的类型就得使用不同的命令
func GetKey(k string) (string, error) {
	val, err := ConnCluster.Get(context.Background(), k).Result()
	if err != nil {
		log.Panicf("k:%v err:%v\n", k, err)
	}
	return val, err
}

func addKeys() {
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("ok_%v", i)
		v := fmt.Sprintf("ov_%v", i)
		ConnCluster.Set(context.Background(), k, v, time.Minute*10)
	}
}

func delKeys() {
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("k_%v", i)
		ConnCluster.Del(context.Background(), k)
		time.Sleep(time.Millisecond)
	}
}

//这段代码说明并发条件下,scan读到的数据不能反映redis当下的真实情况
/*
   整个遍历从开始到结束期间， 一直存在于Redis数据集内的且符合匹配模式的所有Key都会被返回；
   如果发生了rehash，同一个元素可能会被返回多次，遍历过程中新增或者删除的Key可能会被返回，也可能不会。
   from https://tech.meituan.com/2018/07/27/redis-rehash-practice-optimization.html
*/
// func main() {
// 	conn.FlushAll(context.Background())
// 	for i := 0; i < 10; i++ {
// 		k := fmt.Sprintf("k_%v", i)
// 		v := fmt.Sprintf("v_%v", i)
// 		conn.Set(context.Background(), k, v, time.Minute*10)
// 	}
// 	go PrintAllKeys()
// 	// go addKeys()
// 	go delKeys()
// 	time.Sleep(100 * time.Second)
// }

func addZset(key string, member []*redis.Z) error {
	_, err := ConnCluster.ZAdd(context.Background(), key, member...).Result()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
