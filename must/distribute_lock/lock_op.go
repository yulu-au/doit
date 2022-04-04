package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var (
	conn       = newRedisConn()
	delCommand = `if redis.call("GET", KEYS[1]) == ARGV[1] then
    	return redis.call("DEL", KEYS[1])
		else
    	return 0
		end`
)

type redisLock struct {
	ctx    context.Context
	conn   *redis.Client
	key    string
	value  string
	expire time.Duration
}

func (r *redisLock) acquire() (bool, error) {
	ok, err := r.conn.SetNX(r.key, r.value, r.expire).Result()
	if err == redis.Nil {
		//key不存在
		return false, nil
	} else if err != nil {
		//产生了别的错误
		return false, err
	} else if !ok {
		return false, nil
	}
	return true, nil
}

func (r *redisLock) release() (bool, error) {
	val, err := r.conn.Eval(delCommand, []string{r.key}, []string{r.value}).Result()
	if err != nil {
		return false, err
	}
	reply, ok := val.(int64)
	if !ok {
		return false, err
	}
	return reply == 1, nil
}

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
	res := make([]byte, 16)
	for i := 0; i < 16; i++ {
		res[i] = code[rand.Int31n(26)]
	}
	return string(res)
}

func pushOrder(order, val string, wg *sync.WaitGroup) {
	defer wg.Done()
	str := randstr()
	ok, err := conn.SetNX("lockorder", str, time.Second*100).Result()
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
		go pushOrder(v, fmt.Sprintf("order1_detail_%v", i), &wg)
	}
	wg.Wait()
}

func main() {
	conn.FlushAll()
	lockByRedis()
}
