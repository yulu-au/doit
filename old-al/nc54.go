package main

import "sort"

/*
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。

数据范围：0≤n≤30000 \le n \le 30000≤n≤3000，数组中各个元素值满足 ∣val∣≤100|val | \le 100∣val∣≤100
空间复杂度：O(n2)O(n^2)O(n2)，时间复杂度 O(n2)O(n^2)O(n2)

注意：

    三元组（a、b、c）中的元素可以按任意顺序排列。
    解集中不能包含重复的三元组。


*/

func threeSum(num []int) [][]int {
	// write code here
	if len(num) < 3 {
		return [][]int{}
	}
	//对于基础类型的排序
	sort.Ints(num)

	res := make([][]int, 0)
	//定住一个元素,另外两个元素做双指针
	for i := 0; i < len(num)-2; i++ {
		//因为从小到大排序的原因,如果当前大于0,后面不可能小于0
		if num[i] > 0 {
			break
		}
		//去重
		if (i > 0) && (num[i] == num[i-1]) {
			continue
		}
		//窗口一开始是最大的,逐步缩减
		l, r := i+1, len(num)-1
		for l < r {
			if 0 == num[i]+num[l]+num[r] {
				res = append(res, []int{num[i], num[l], num[r]})
				//相等的话,两边同时加减有可能还等的,另外需要去重
				lt, rt := l, r
				for num[lt] == num[l] && lt < r {
					lt++
				}
				for num[rt] == num[r] && rt > l {
					rt--
				}
				l = lt
				r = rt
			} else if 0 < num[i]+num[l]+num[r] {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

/*
func threeSum(nums []int) [][]int {
  count:=len(nums)
  if count<3{
      return [][]int{}
  }

  sort.Ints(nums) // 排序后简化的问题，可以使用滑动窗口
  ans := make([][]int, 0)
  // 寻找 -a=b+c
  for a:=0;a<count;a++{
      if a>0 && (nums[a]==nums[a-1]){ //去掉重复a，因为排过序，所以只跟上一个比
          continue
      }
      target:=(-nums[a])
      c:=count-1 //如果c=b+1可能会越界
      for b:=a+1;b<count;b++{ // 固定b找c
          if (b>a+1) && nums[b]==nums[b-1]{ //去掉重复b
              continue
          }
          for b<c && (nums[b]+ nums[c]>target) {
              c--
          }
          if b==c {
              break
          }
          if nums[b]+ nums[c]==target{
              ans=append(ans,[]int{nums[a],nums[b],nums[c]})
          }

      }
  }
  return ans
}

*/
