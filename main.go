package main

import (
	"doit/must/doredis"
	"fmt"
	"time"
)

func main() {
	res := doredis.Top10("rank-course")
	for _, v := range res {
		fmt.Println(v)
	}
	time.Sleep(time.Second)
	doredis.AddScore("rank-course", -99, "course-69")
	fmt.Println("------------------")
	res = doredis.Top10("rank-course")
	for _, v := range res {
		fmt.Println(v)
	}
}
