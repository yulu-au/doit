package main

/*
 给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）

数据范围：0≤n≤15000 \le n \le 15000≤n≤1500,树上每个节点的val满足 ∣val∣<=1500|val| <= 1500∣val∣<=1500
要求：空间复杂度：O(n)O(n) O(n)，时间复杂度：O(n)O(n)O(n)
例如：
给定的二叉树是{1,2,3,#,#,4,5}
*/

func Print(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	qe := make([]*TreeNode, 0)
	qe = append(qe, pRoot)
	//标志 用来看是不是翻转
	fg := false
	for len(qe) != 0 {
		ln := len(qe)
		tmp := make([]int, 0)
		for i := 0; i < ln; i++ {
			cur := qe[0]
			qe = qe[1:]

			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				qe = append(qe, cur.Left)
			}
			if cur.Right != nil {
				qe = append(qe, cur.Right)
			}

		}
		if fg {
			reverse(tmp)
		}
		fg = !fg
		res = append(res, tmp)
	}
	return res
}

func reverse(list []int) {
	lt, rt := 0, len(list)-1
	for lt < rt {
		list[lt], list[rt] = list[rt], list[lt]
		lt++
		rt--
	}
}
