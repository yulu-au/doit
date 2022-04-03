//收集好看的博客
//https://www.lixueduan.com/categories/Redis/
package main

import (
	"context"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

var (
	conn = newRedisConn()
)

func newRedisConn() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
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

func scanAllKeys() {
	all := make([]string, 0)
	var cur uint64 = 0
	for {
		time.Sleep(time.Millisecond)
		res, cursor, err := conn.Scan(context.Background(), cur, "*", 5).Result()
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

func addKeys() {
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("ok_%v", i)
		v := fmt.Sprintf("ov_%v", i)
		conn.Set(context.Background(), k, v, time.Minute*10)
	}
}

func delKeys() {
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("k_%v", i)
		conn.Del(context.Background(), k)
		time.Sleep(time.Millisecond)
	}
}

//这段代码说明并发条件下,scan读到的数据不能反映redis当下的真实情况
/*
   整个遍历从开始到结束期间， 一直存在于Redis数据集内的且符合匹配模式的所有Key都会被返回；
   如果发生了rehash，同一个元素可能会被返回多次，遍历过程中新增或者删除的Key可能会被返回，也可能不会。
   from https://tech.meituan.com/2018/07/27/redis-rehash-practice-optimization.html
*/
func main() {
	conn.FlushAll(context.Background())
	for i := 0; i < 10; i++ {
		k := fmt.Sprintf("k_%v", i)
		v := fmt.Sprintf("v_%v", i)
		conn.Set(context.Background(), k, v, time.Minute*10)
	}
	go scanAllKeys()
	// go addKeys()
	go delKeys()

	time.Sleep(100 * time.Second)
}
