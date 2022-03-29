package main

/*
 实现函数 int sqrt(int x).
计算并返回 x 的平方根（向下取整）

数据范围： 0<=x<231−10 <= x < 2^{31}-10<=x<231−1
要求：空间复杂度 O(1)O(1)O(1)，时间复杂度 O(logx)O(logx)O(logx)
*/

func sqrt(x int) int {
	//定义红蓝区间,红区间指的是平方数小于或者等于x的,蓝区间是大于的
	//l r是红蓝区间的端点,初始时红蓝区间是空的
	l, r := -1, x+1
	//运行条件是二者不相邻
	for l < r-1 {
		m := l + (r-l)>>1
		tmp := m * m
		//说明m落在红区间里
		if tmp <= x {
			l = m
			//说明m落在蓝区间里
		} else {
			r = m
		}
	}
	//红区间的最右边就是所求
	return l
}
