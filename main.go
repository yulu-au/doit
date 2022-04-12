package main

import (
	"doit/must/doredis"
	"fmt"
)

func main() {
	sliK := doredis.PrintAllKeysFromCluster()
	// fmt.Println(sliK)
	for _, v := range sliK {
		r, _ := doredis.GetKey(v)
		fmt.Printf("get : %v\n", r)
	}
}
