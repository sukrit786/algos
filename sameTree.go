/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	s1 := []*TreeNode{}
	s2 := []*TreeNode{}

	s1 = append(s1, p)
	s2 = append(s2, q)

	for len(s1) > 0 && len(s2) > 0 {
		top1, top2 := s1[0], s2[0]

		// if top1.Val!=top2.Val {
		//     return false
		// }
		s1, s2 = s1[1:], s2[1:]

		if top1.Left != nil && top2.Left == nil || top1.Right != nil && top2.Right == nil || top2.Left != nil && top1.Left == nil || top2.Right != nil && top1.Right == nil {
			return false
		}

		if top1.Val != top2.Val {
			return false
		}

		if top1.Left != nil {
			s1 = append(s1, top1.Left)
			s2 = append(s2, top2.Left)
		}
		if top1.Right != nil {
			s1 = append(s1, top1.Right)
			s2 = append(s2, top2.Right)
		}

	}

	fmt.Print(s1, s2)
	return true
}

//just check preorder
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}