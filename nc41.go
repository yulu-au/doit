package main

/*
 给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

数据范围：0≤arr.length≤1050\le arr.length \le 10^50≤arr.length≤105，0<arr[i]≤1050 < arr[i] \le 10^50<arr[i]≤105
*/

func maxLength(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	set := make(map[int]int)
	ret := 0

	i, cnt := 0, 0

	for i < len(arr) {
		val := arr[i]
		if _, exist := set[val]; exist {
			delete(set, val)
			set[val] = i
			cnt = len(set)
		} else {
			set[val] = i
			cnt = len(set)
		}
		i++
		ret = max(ret, cnt)
	}

	return ret
}
