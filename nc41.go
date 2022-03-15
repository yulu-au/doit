package main

/*
 给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

数据范围：0≤arr.length≤1050\le arr.length \le 10^50≤arr.length≤105，0<arr[i]≤1050 < arr[i] \le 10^50<arr[i]≤105
*/

//滑动窗口,使用map的key做唯一性约束,val存数组索引位置
func maxLength(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	set := make(map[int]int)
	ret := 0

	i, j, cnt := 0, 0, 0

	for i < len(arr) {
		val := arr[i]
		if index, exist := set[val]; exist {
			//注意这里
			j = max(j, index+1) //index+1 may be number smaller than j
			delete(set, val)
			set[val] = i
			cnt = i - j + 1
		} else {
			set[val] = i
			cnt = i - j + 1
		}
		i++
		ret = max(ret, cnt)
	}

	return ret
}

/*

func maxLength( arr []int ) int {
  // write code here
  if len(arr)==0 {
      return 0
  }
  max,start:=0,0
  h:=map[int]int{}
  for end:=0;end<len(arr);end++{
      if h[arr[end]]!=0 {
         start=Max(start,h[arr[end]]+1) //历史最大值，因为可能新的重复索引要小于start，更新成重复的索引下一个
      }
       h[arr[end]]=end
       max=Max(max,end-start+1)
  }
  return max
}

*/
