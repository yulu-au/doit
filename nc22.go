package main

/*
给出一个有序的整数数组 A 和有序的整数数组 B ，
请将数组 B 合并到数组 A 中，变成一个有序的升序数组
*/

func merge(A []int, m int, B []int, n int) {
	if n == 0 {
		return
	}
	a, b := iterator{A, 0, m - 1}, iterator{B, 0, n - 1}
	if m == 0 {
		for k := m + n - 1; k >= 0; k-- {
			A[k] = b.returnMax()
		}
		return
	}
	for k := m + n - 1; k >= 0; k-- {
		va, oka := a.max()
		vb, okb := b.max()
		if !oka {
			A[k] = b.returnMax()
			continue
		}
		if !okb {
			A[k] = a.returnMax()
			continue
		}

		if va > vb {
			A[k] = a.returnMax()
		} else {
			A[k] = b.returnMax()
		}
	}

	// write code here
}

type iterator struct {
	sli []int
	lt  int
	rt  int
}

func (i *iterator) isAsc() bool {
	if len(i.sli) == 1 {
		return true
	}
	if i.sli[0] > i.sli[1] {
		return false
	}
	return true
}

func (i *iterator) max() (int, bool) {
	if i.lt > i.rt {
		return -1, false
	}
	if i.isAsc() {
		r := i.sli[i.rt]
		return r, true
	} else {
		r := i.sli[0]
		return r, true
	}
}

func (i *iterator) returnMax() int {
	if i.isAsc() {
		r := i.sli[i.rt]
		i.rt--
		return r
	} else {
		r := i.sli[0]
		i.lt++
		return r
	}
}
