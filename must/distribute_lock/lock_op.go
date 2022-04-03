package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/go-redis/redis"
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

func randstr() string {
	code := []byte("abcdefghigklmnopqrstuvwxyz")
	res := make([]byte, 0)
	for i := 0; i < 26; i++ {
		res = append(res, code[rand.Int31n(26)])
	}
	return string(res)
}

func pushOrder(order, val string, wg *sync.WaitGroup) {
	defer wg.Done()
	str := randstr()
	ok, err := conn.SetNX("lockorder", str, 0).Result()
	if !ok {
		fmt.Println("already do order1")
		return
	}

	err = conn.LPush(order, val).Err()
	if err != nil {
		panic(err)
	}
}

func lockByRedis() {
	sliOrder := []string{"order1", "order1", "order1"}
	var wg sync.WaitGroup
	wg.Add(3)
	//三个goroutine模拟三个节点想要插入订单列表
	//理想情况应该只有一个插入才行
	for i, v := range sliOrder {
		go pushOrder(v, fmt.Sprintf("val_%v", i), &wg)
	}
	wg.Wait()
}

func main() {
	conn.FlushAll()
	lockByRedis()
}
