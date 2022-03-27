package main

/*
 给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

数据范围：树上节点数满足 1≤n≤105 1 \le n \le 10^5 \ 1≤n≤105  , 节点值val满足区间 [0,n)
要求：时间复杂度 O(n)O(n)O(n)

注：本题保证二叉树中每个节点的val值均不相同。
*/

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	//存储子节点到父节点的映射
	M := make(map[int]int, 0)
	r1, r2 := make([]int, 0), make([]int, 0)
	x, y := o1, o2

	//层序遍历
	M[root.Val] = -1
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Left != nil {
			//记录映射
			M[cur.Left.Val] = cur.Val
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			M[cur.Right.Val] = cur.Val
			queue = append(queue, cur.Right)
		}
	}
	//子节点到根节点的路径
	r1 = append(r1, x)
	for {
		if v := M[x]; v == -1 {
			break
		} else {
			r1 = append(r1, v)
			x = v
		}

	}
	//路径
	r2 = append(r2, y)
	for {
		if v := M[y]; v == -1 {
			break
		} else {
			r2 = append(r2, v)
			y = v
		}
	}
	//从后向前寻找第一个不一样的节点,它前面就是所求
	index1, index2 := len(r1)-1, len(r2)-1
	for index1 >= 0 && index2 >= 0 {
		if r1[index1] != r2[index2] {
			return r1[index1+1]
		}
		index1--
		index2--
	}
	//逻辑到这里必然是有一个切片遍历完了,两个切片是包含与被包含的关系
	//r1走完全程
	if index1 == -1 {
		return r1[0]
	}
	//或者r2走完全程
	return r2[0]
}
