package main

/*
 给出一个整型数组 numbers 和一个目标值 target，请在数组中找出两个加起来等于目标值的数的下标，返回的下标按升序排列。
（注：返回的数组下标从1开始算起，保证target一定可以由数组里面2个数字相加得到）

数据范围：2≤len(numbers)≤1052\leq len(numbers) \leq 10^52≤len(numbers)≤105，−10≤numbersi≤109-10 \leq numbers_i \leq 10^9−10≤numbersi​≤109，0≤target≤1090 \leq target \leq 10^90≤target≤109
要求：空间复杂度 O(n)O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)O(nlogn)
*/

func twoSum(numbers []int, target int) []int {
	store := make(map[int]int, 0)

	for i, v := range numbers {
		if vv, exist := store[target-v]; exist {
			if i > vv {
				return []int{vv + 1, i + 1}
			}
			return []int{i + 1, vv + 1}
		} else {
			store[v] = i
		}
	}
	return []int{}
}
