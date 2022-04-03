package main

/*
 给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。
数据范围：0≤k,n≤100000\le k,n \le 100000≤k,n≤10000，数组中每个数的大小0≤val≤10000 \le val \le 1000 0≤val≤1000
要求：空间复杂度 O(n)O(n)O(n) ，时间复杂度 O(nlogn)O(nlogn)O(nlogn)
*/

func GetLeastNumbers_Solution(input []int, k int) []int {
	res := make([]int, 0)
	h := make(intHeap, 0)
	ph := &h
	for _, v := range input {
		ph.Push(v)
	}

	for i := 0; i < k; i++ {
		r := ph.Pop()
		res = append(res, r)
	}

	return res
}

type intHeap []int

func (p *intHeap) Init() {
	if p.Len() == 0 {
		return
	}

	for i := p.Len() - 1; i >= 0; i-- {
		p.down(i)
	}
}

func (in intHeap) Less(i, j int) bool {
	return in[i] < in[j]
}

func (in intHeap) Len() int {
	return len(in)
}

func (p *intHeap) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *intHeap) Pop() int {
	x := (*p)[0]
	p.Swap(0, p.Len()-1)
	*p = (*p)[:len(*p)-1]
	p.down(0)
	return x
}

func (p *intHeap) Push(x int) {
	*p = append(*p, x)
	p.up(p.Len() - 1)
}

func (p *intHeap) up(i int) {
	x := (i - 1) / 2
	if x < 0 {
		return
	}

	for p.Less(i, x) {
		p.Swap(i, x)

		i = x
		x = (x - 1) / 2
		if x < 0 {
			break
		}
	}
}

func (p *intHeap) down(i int) {
	child := p.theSmallerChild(i)
	if child > p.Len()-1 || child == -1 {
		return
	}

	for !p.Less(i, child) {
		if child+1 > p.Len()-1 {

		} else {
			if !p.Less(child, child+1) {
				child++
			}
		}
		p.Swap(i, child)
		i = child
		child = p.theSmallerChild(i)
		if child == -1 {
			break
		}
	}
}

func (in intHeap) theSmallerChild(root int) int {
	a := root*2 + 1
	b := a + 1
	ln := in.Len() - 1
	if a > ln && b > ln {
		return -1
	}
	if a > ln {
		return b
	}
	if b > ln {
		return a
	}
	if in[a] > in[b] {
		return b
	} else {
		return a
	}
}
