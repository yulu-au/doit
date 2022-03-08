package main

/*
有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
给定一个整数数组 a ,同时给定它的大小n和要找的 k ，
请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。
要求：时间复杂度 O(nlogn)O(nlogn)O(nlogn)，空间复杂度 O(1)O(1)O(1)
数据范围：0≤n≤10000\le n \le 10000≤n≤1000， 1≤K≤n1 \le K \le n1≤K≤n，
数组中每个元素满足 0≤val≤100000000 \le val \le 100000000≤val≤10000000
*/

func findKth(arr []int, n, k int) int {
	if len(arr) == 0 {
		return -1
	}
	l, r := 0, len(arr)-1
	//partition is a list (2,4,5,5,7)
	kindex := n - k
	for {
		i := partition(l, r, arr)
		if i == kindex {
			break
		}
		if i < kindex {
			l = i + 1
			// kindex = kindex - i - 1
		} else {
			r = i - 1

		}
	}
	return arr[kindex]
}
