package main

/*
述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。
求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。

数据范围：1≤n≤401 \leq n \leq 401≤n≤40
要求：时间复杂度：O(n)O(n)O(n) ，空间复杂度： O(1)O(1)O(1)
*/
// func jumpFloor(number int) int {
// 	if number == 1 {
// 		return 1
// 	}
// 	if number == 2 {
// 		return 2
// 	}

// 	return jumpFloor(number) + jumpFloor(number-1)
// }
func jumpFloor(number int) int {
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}

	a, b := 1, 2
	x := 0
	for i := 3; i <= number; i++ {
		x = a + b
		a = b
		b = x
	}

	return x
}
