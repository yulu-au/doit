package main

/*
 给定一棵二叉树，分别按照二叉树先序，中序和后序打印所有的节点
*/

func threeOrders(root *TreeNode) [][]int {
	res := [][]int{}
	res = append(res, PreOrderTraversal2(root), InOrderTraversal2(root), PostOrderTraversal2(root))
	return res
}

func PreOrderTraversal(root *TreeNode) []int {
	res := []int{}
	preOrderTravsal(root, &res)
	return res
}

func preOrderTravsal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrderTravsal(root.Left, res)
	preOrderTravsal(root.Right, res)
}

func InOrderTravsal(root *TreeNode) []int {
	res := []int{}
	inOrderTravsal(root, &res)
	return res
}

func inOrderTravsal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrderTravsal(root.Left, res)
	*res = append(*res, root.Val)
	inOrderTravsal(root.Right, res)
}

func PostOrderTravsal(root *TreeNode) []int {
	res := []int{}
	postOrderTravsal(root, &res)
	return res
}

func postOrderTravsal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postOrderTravsal(root.Left, res)
	postOrderTravsal(root.Right, res)
	*res = append(*res, root.Val)
}

//
func PreOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		first := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if first == nil {
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res, r.Val)
		} else {
			// stack = append(stack, first.Right, first.Left, nil, first)
			if first.Right != nil {
				stack = append(stack, first.Right)
			}
			if first.Left != nil {
				stack = append(stack, first.Left)
			}
			stack = append(stack, first, nil)
		}
	}
	return res
}

func InOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		first := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if first == nil {
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res, r.Val)
		} else {
			// stack = append(stack, first.Right, first.Left, nil, first)
			if first.Right != nil {
				stack = append(stack, first.Right)
			}

			stack = append(stack, first, nil)

			if first.Left != nil {
				stack = append(stack, first.Left)
			}
		}
	}
	return res
}

func PostOrderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		first := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if first == nil {
			r := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res = append(res, r.Val)
		} else {
			// stack = append(stack, first.Right, first.Left, nil, first)
			stack = append(stack, first, nil)

			if first.Right != nil {
				stack = append(stack, first.Right)
			}
			if first.Left != nil {
				stack = append(stack, first.Left)
			}
		}
	}
	return res
}
