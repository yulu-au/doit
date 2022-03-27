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
		//说明m落在了在左区间
		if rotateArray[m] > rotateArray[r] {
			l = m
			//说明m落在了在右区间
		} else if rotateArray[m] < rotateArray[r] {
			r = m
			//不确定m在哪里,端点可以左移一步
		} else {
			r--
		}
	}
	return rotateArray[r]
}

// func minNumberInRotateArray( rotateArray []int ) int {
//     // write code here
//     left, right := 0, len(rotateArray)-1

//     for left < right {
//         middle := left + (right - left)>>1
//         //如果中间值大于最右边的值，说明旋转之后最小的
//         //数字肯定在mid的右边，比如[3, 4, 5, 6, 7, 1, 2]
//         if rotateArray[middle] > rotateArray[right] {
//             left = middle + 1
//         }else if rotateArray[middle] < rotateArray[right] {
//             //如果中间值小于最右边的值，说明旋转之后最小的
//             //数字肯定在mid的前面，比如[6, 7, 1, 2, 3, 4, 5],
//             //注意这里mid是不能减1的，比如[3，1，3]，我们这里只是
//             //证明了numbers[mid]比numbers[right]小，但有可能
//             //numbers[mid]是最小的，所以我们不能把它给排除掉
//             right = middle
//         }else {
//             //如果中间值等于最后一个元素的值，我们是没法确定最小值是
//              // 在mid的前面还是后面，但我们可以缩小查找范围，让right
//              // 减1，因为即使right指向的是最小值，但因为他的值和mid
//              // 指向的一样，我们这里并没有排除mid，所以结果是不会有影响的。
//              //比如[3，1，3，3，3，3，3]和[3，3，3，3，3，1，3],中间的值
//              //等于最右边的值，但我们没法确定最小值是在左边还是右边
//             right--
//         }
//     }

// //  不断的缩小查找范围，当查找范围的长度为1的时候返回，也就是left等于right的时候
//     return rotateArray[left]
// }
