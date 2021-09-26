/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func isSymmetric(root *TreeNode) bool {
//     if root == nil {
//         return true
//     }

//     return checkSymmetric(root.Left,root.Right)

// }

// func checkSymmetric(left *TreeNode, right *TreeNode) bool {
//     if left == nil || right == nil {
//         return left == right
//     }

//     if left.Val != right.Val {
//         return false
//     }

//     return checkSymmetric(left.Left,right.Right) && checkSymmetric(left.Right,right.Left)
// }
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i] == nil && queue[size-i-1] == nil {
				continue
			}
			if queue[i] == nil || queue[size-i-1] == nil {
				return false
			}
			if queue[i].Val != queue[size-i-1].Val {
				return false
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		queue = queue[size:]
	}
	return true
}


