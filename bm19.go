package main

/*
 给定一个长度为n的数组nums，请你找到峰值并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个所在位置即可。
1.峰值元素是指其值严格大于左右相邻值的元素。严格大于即不能有等于
2.假设 nums[-1] = nums[n] = −∞-\infty−∞
3.对于所有有效的 i 都有 nums[i] != nums[i + 1]
4.你可以使用O(logN)的时间复杂度实现此问题吗？

数据范围：
1≤nums.length≤2×105 1 \le nums.length \le 2\times 10^5 \ 1≤nums.length≤2×105
−231<=nums[i]<=231−1-2^{31}<= nums[i] <= 2^{31} - 1−231<=nums[i]<=231−1

*/

/*
红蓝区间模板方案
*/
func findPeakElement(nums []int) int {
	//l r 初始都是指向不存在的位置
	l, r := -1, len(nums)
	//循环条件
	for l+1 != r {
		//m的取值范围是0到r
		m := l + (r-l)>>1
		if isBlue(nums, m) {
			l = m
		} else {
			r = m
		}
	}
	//返回值是l或者r
	return r
}

func isBlue(arr []int, m int) bool {
	if m+1 > len(arr)-1 {
		return false
	}
	if arr[m] < arr[m+1] {
		return true
	}
	return false
}
