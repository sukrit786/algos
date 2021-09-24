package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

func (root *tree) levelorder() {
	q := []*tree{}
	q = append(q, root)
	if root == nil {
		return
	}
	for len(q) != 0 {
		top := q[0]
		fmt.Print(top.val, " ")
		if top.left != nil {
			q = append(q, top.left)
		}
		if top.right != nil {
			q = append(q, top.right)
		}
		q = q[1:]
	}
}

func (root *tree) inorder() {
	if root == nil {
		return
	}
	root.left.inorder()
	fmt.Print(root.val, " ")
	root.right.inorder()
}

//iterative Inorder
func inorderTraversal(root *TreeNode) []int {
	s := []*TreeNode{}
	ans := []int{}

	node := root
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		node = s[len(s)-1]
		s = s[:len(s)-1]
		ans = append(ans, node.Val)
		node = node.Right
	}
	return ans

}

func insert(root *tree, arr []int, i int, n int) *tree {
	if i < n {
		temp := &tree{val: arr[i]}
		root = temp
		root.left = insert(root.left, arr, 2*i+1, n)
		root.right = insert(root.right, arr, 2*i+2, n)
	}
	return root
}
func main() {
	// your code goes here
	arr := []int{1, 2, 3, 4, 5, 6}
	root := insert(&tree{}, arr, 0, len(arr))
	fmt.Println("i was called by baka ", root)
	root.levelorder()
}
