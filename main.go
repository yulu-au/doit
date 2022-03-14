package main

func main() {
	a := make([]int, 0, 6)
	a = append(a, 4, 5, 6, 0, 0, 0)
	merge(a, 3, []int{1, 2, 3}, 3)
}
