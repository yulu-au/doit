package must

import "fmt"

//问题出在v上,因为数组的v是不会变的
func Bug01() {
	a, b := [3]int{1, 2, 3}, []int{1, 2, 3}

	for i, v := range a {
		if i == 0 {
			a[0], a[1] = 100, 200
		}
		a[i] = 100 + v
	}
	//101,102,103
	fmt.Println(a)

	for i, v := range b {
		if i == 0 {
			if i == 0 {
				b[0], b[1] = 100, 200
			}
		}
		b[i] = 100 + v
	}
	//101,300,103
	fmt.Println(b)
}
