package main

import (
	"context"
	"fmt"
	"log"
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
	race int //race代表某个不在这个进程内的资源,需要加分布式锁访问
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

func doRace(wg *sync.WaitGroup, task string) {
	defer wg.Done()

	for {
		//初始化分布式锁 redislock
		rl := redisLock{
			ctx:    context.Background(),
			conn:   conn,
			key:    "placeholder",
			value:  randstr(),
			expire: time.Second * 100,
		}
		ok, err := rl.acquire()
		if err != nil {
			log.Println(err)
		}
		if !ok {
			continue
		}

		defer rl.release()

		//这里开始
		race++
		return
	}
}

func lockByRedis(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	//三个goroutine同时修改race这个key
	for i := 0; i < n; i++ {
		go doRace(&wg, "task")
	}
	wg.Wait()
}

func main() {
	conn.FlushAll()
	//修改被争用的资源 race
	lockByRedis(1000)
	fmt.Printf("race : %v\n", race)
}
