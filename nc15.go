package main

/*
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
例如：
给定的二叉树是{3,9,20,#,#,15,7},


该二叉树层序遍历的结果是
[
[3],
[9,20],
[15,7]
]
*/
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	workQuene := make([]*TreeNode, 0)
	if root == nil {
		return res
	}

	workQuene = append(workQuene, root)
	for len(workQuene) != 0 {
		//how many node this loop do
		ln := len(workQuene)
		tmp := make([]int, 0)
		for i := 0; i < ln; i++ {
			cur := workQuene[0]
			workQuene = workQuene[1:]
			tmp = append(tmp, cur.Val)

			if cur.Left != nil {
				workQuene = append(workQuene, cur.Left)
			}
			if cur.Right != nil {
				workQuene = append(workQuene, cur.Right)
			}
		}

		res = append(res, tmp)
	}

	return res
}
