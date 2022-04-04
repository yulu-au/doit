package language

import "fmt"

//清空切片的几种方法
func howSlice() {
	origin := [...]int{1, 2, 3, 4, 5}
	sli := origin[:4]
	fmt.Printf("%p  %p %p %p\n", &origin, &origin[0], sli, &sli[0])

	//切片指向的底层数组没变,只是相当于清空了
	sli = sli[0:0]
	fmt.Printf("%p  %p %p\n", &origin, &origin[0], sli)
	//0 5
	fmt.Printf("%v %v\n", len(sli), cap(sli))

	//切片这个引用类型被赋值nil
	sli = nil
	//切片指向的底层数组改了
	sli = []int{}
}
