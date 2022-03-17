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

	l, r := -1, len(rotateArray)
	for l+1 != r {
		m := l + (r-l)>>1
		if rotateArray[m] >= rotateArray[0] && rotateArray[m] != rotateArray[len(rotateArray)-1] {
			l = m
		} else {
			r = m
		}
	}
	return rotateArray[r]
}
