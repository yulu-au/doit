package main

/*
有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，
即把一个数组最开始的若干个元素搬到数组的末尾，变成一个旋转数组，比如变成了[3,4,5,1,2]，
或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。
*/
//细节还需打磨
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return -1
	}
	/*
		根据红蓝区间模板,以及题意,端点其实都一定是在区间里
	*/
	l, r := 0, len(rotateArray)-1
	for l+1 != r {
		m := l + (r-l)>>1
		//说明m一定在左区间
		if rotateArray[m] > rotateArray[r] {
			l = m
			//说明m一定在右区间
		} else if rotateArray[m] < rotateArray[r] {
			r = m
			//不确定m在哪里,不过右区间断点可以左移一步
		} else {
			r--
		}
	}
	return rotateArray[r]
}
