package main

import "fmt"

type node struct {
	val   int
	left  *node
	right *node
}

func (root *node) insert(data int) {
	el := &node{val: data}
	if root.val > data {
		if root.left == nil {
			root.left = el
		} else {
			root.left.insert(data)
		}
	} else {
		if root.right == nil {
			root.right = el
		} else {
			root.right.insert(data)
		}
	}
}
func (root *node) inorder() {
	if root == nil {
		return
	}
	root.left.inorder()
	fmt.Print(root.val, " ")
	root.right.inorder()
}
func (root *node) preorder() {
	if root == nil {
		return
	}
	fmt.Print(root.val, " ")
	root.left.inorder()
	root.right.inorder()
}
func (root *node) postorder() {
	if root == nil {
		return
	}
	root.left.inorder()
	root.right.inorder()
	fmt.Print(root.val, " ")
}

//   while (q.empty() == false)
//     {
//         // Print front of queue and remove it from queue
//         Node *node = q.front();
//         cout << node->data << " ";
//         q.pop();

//         /* Enqueue left child */
//         if (node->left != NULL)
//             q.push(node->left);

//         /*Enqueue right child */
//         if (node->right != NULL)
//             q.push(node->right);
//     }
func (root *node) levelorder() {
	q := []*node{}
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
		// 		fmt.Println(q)
	}
}

func main() {
	root := &node{val: 50}
	root.insert(30)
	root.insert(15)
	root.insert(20)
	root.insert(10)
	root.insert(40)
	root.insert(60)
	fmt.Print("InOrder => ")
	root.inorder()
	fmt.Println()
	fmt.Print("PreOrder => ")
	root.preorder()
	fmt.Println()
	fmt.Print("PostOrder => ")
	root.postorder()
	fmt.Println()
	fmt.Print("LevelOrder => ")
	root.levelorder()
}
