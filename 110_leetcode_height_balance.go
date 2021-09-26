/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// func isBalanced(root *TreeNode) bool {
//     if root == nil {
//         return true
//     }
//     var height func(*TreeNode) int

//     height = func(root *TreeNode) int {
//         if root == nil {
//             return 0
//         }

//         return 1+max(height(root.Left),height(root.Right))

//     }
//     xx := height(root.Left)
//     yy := height(root.Right)
//     fmt.Print(xx,yy)
//     diff:=(math.Abs(float64(xx)-float64(yy)))
//     if(diff>1) {
//         return false
//     }
//     return isBalanced(root.Left) && isBalanced(root.Right)
// }
func absDiff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(*TreeNode) int

	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftHeight := height(root.Left)
		rightHeight := height(root.Right)
		bf := absDiff(leftHeight, rightHeight)
		if bf > 1 || leftHeight == -1 || rightHeight == -1 {
			return -1
		}

		return 1 + max(height(root.Left), height(root.Right))

	}
	return height(root) != -1
}

// func isBalanced(root *TreeNode) bool {
// 	if root == nil {
// 		return true
// 	}

// 	lh := height(root.Left)
// 	rh := height(root.Right)

// 	if absDiff(lh, rh) > 1 {
// 		return false
// 	}

// 	return isBalanced(root.Left) && isBalanced(root.Right)
// }

// func absDiff(a, b int) int {
// 	if a >= b {
// 		return a - b
// 	}
// 	return b - a
// }

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + maxHeight(height(root.Left), height(root.Right))
}

// func maxHeight(a, b int) int {
// 	if a >= b {
// 		return a
// 	}
// 	return b
// }