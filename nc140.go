package main

func MySort(arr []int) []int {
	quickSort2(0, len(arr)-1, arr)
	return arr
}

func quickSort(left, right int, arr []int) {
	if left >= right {
		return
	}

	index := partition(left, right, arr)
	quickSort(left, index-1, arr)
	quickSort(index+1, right, arr)
}

func partition(left, right int, arr []int) int {
	code := arr[left]
	for left < right {
		for arr[right] > code && left < right {
			right--
		}
		arr[left] = arr[right]
		for arr[left] <= code && left < right {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = code
	return left
}

//
func quickSort2(left, right int, arr []int) {
	stack := make([][]int, 0)
	stack = append(stack, []int{left, right})
	for len(stack) != 0 {
		cur := stack[0]
		stack = stack[1:]

		mid := partition(cur[0], cur[1], arr)
		if (mid - 1) > cur[0] {
			stack = append(stack, []int{cur[0], mid - 1})
		}
		if (mid + 1) < cur[1] {
			stack = append(stack, []int{mid + 1, cur[1]})
		}
	}
}
