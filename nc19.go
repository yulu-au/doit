package main

/*
 输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，
 子数组最小长度为1。求所有子数组的和的最大值
*/
/*
定义dp[i]表示前i个元素(must have ith node)的连续子数组的最大和


dp[i] = max(dp[i-1]+array[i],array[i])
*/

func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	if len(array) == 0 {
		return -1
	}
	dp := array[0]
	ret := dp
	array = array[1:]
	for _, v := range array {
		dp = max(v+dp, v)
		ret = max(dp, ret)
	}
	return ret
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}