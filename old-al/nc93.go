package main

import (
	"container/list"
)

/*
描述
设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为 k ，并有如下两个功能
1. set(key, value)：将记录(key, value)插入该结构
2. get(key)：返回key对应的value值

提示:
1.某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的，然后都会刷新缓存。
2.当缓存的大小超过k时，移除最不经常使用的记录。
3.输入一个二维数组与k，二维数组每一维有2个或者3个数字，第1个数字为opt，第2，3个数字为key，value
若opt=1，接下来两个整数key, value，表示set(key, value)
若opt=2，接下来一个整数key，表示get(key)，若key未出现过或已被移除，则返回-1
对于每个opt=2，输出一个答案
4.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹

要求：set和get操作复杂度均为 O(1)O(1)O(1)
示例1
输入：

[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3

返回值：

[1,-1]

说明：

[1,1,1]，第一个1表示opt=1，要set(1,1)，即将(1,1)插入缓存，缓存是{"1"=1}
[1,2,2]，第一个1表示opt=1，要set(2,2)，即将(2,2)插入缓存，缓存是{"1"=1,"2"=2}
[1,3,2]，第一个1表示opt=1，要set(3,2)，即将(3,2)插入缓存，缓存是{"1"=1,"2"=2,"3"=2}
[2,1]，第一个2表示opt=2，要get(1)，返回是[1]，因为get(1)操作，缓存更新，缓存是{"2"=2,"3"=2,"1"=1}
[1,4,4]，第一个1表示opt=1，要set(4,4)，即将(4,4)插入缓存，但是缓存已经达到最大容量3，移除最不经常使用的{"2"=2}，插入{"4"=4}，缓存是{"3"=2,"1"=1,"4"=4}
[2,2]，第一个2表示opt=2，要get(2)，查找不到，返回是[1,-1]

示例2
输入：

[[1,1,1],[1,2,2],[2,1],[1,3,3],[2,2],[1,4,4],[2,1],[2,3],[2,4]],2

返回值：

[1,-1,-1,3,4]

备注：

1≤K≤N≤1051 \leq K \leq N \leq 10^51≤K≤N≤105
−2×109≤x,y≤2×109-2 \times 10^9 \leq x,y \leq 2 \times 10^9−2×109≤x,y≤2×109
*/
func LRU(op [][]int, k int) []int {
	var c cache
	c.Init(k)
	res := make([]int, 0)

	for _, v := range op {
		if v[0] == 1 {
			c.Set(v[1], v[2])
		} else {
			r, _ := c.Get(v[1])
			res = append(res, r)
		}
	}
	return res
}

type lruVal struct {
	ptr *list.Element
	val int
}

//get set O(1)复杂度 靠的是map结构
//list+cap用来限制缓存的大小,丢弃策略
type cache struct {
	M    map[int]lruVal
	List list.List
	cap  int
}

func (c *cache) Init(k int) {
	c.M = make(map[int]lruVal)
	c.List.Init()
	c.cap = k
}

func (c *cache) Get(key int) (int, bool) {
	if c.have(key) {
		c.List.MoveToFront(c.M[key].ptr)
		return c.M[key].val, true
	}
	return -1, false
}

func (c *cache) Set(key, val int) bool {
	if c.have(key) {
		c.List.Remove(c.M[key].ptr)
	}
	c.insert(key, val)

	if c.List.Len() > c.cap {
		last := c.List.Back()
		v := last.Value.(int)
		delete(c.M, v)
		c.List.Remove(last)
	}
	return true
}

func (c *cache) have(key int) bool {
	_, ok := c.M[key]
	if ok {
		return true
	}
	return false
}

func (c *cache) insert(key, val int) {
	p := c.List.PushFront(key)
	v := lruVal{p, val}
	c.M[key] = v
}

// func (c *cache) remove(key int) {
// 	c.List.Remove(c.M[key].ptr)
// }
