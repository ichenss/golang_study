package TreeNode

import "fmt"

// TreeNode 二叉树节点定义
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

var Root *TreeNode

func Insert(nums []int) {
	Root = &TreeNode{val: nums[0]}
	for i := 1; i < len(nums); i++ {
		p := Root
		newNode := &TreeNode{val: nums[i]}
		for {
			if p.val < nums[i] {
				if p.right == nil {
					p.right = newNode
					break
				}
				p = p.right
			} else if p.val > nums[i] {
				if p.left == nil {
					p.left = newNode
					break
				}
				p = p.left
			} else {
				fmt.Println("Value already exists in the tree")
				return
			}
		}
	}
}

// Inorder 中序遍历递归实现
func Inorder(root *TreeNode) {
	if root == nil {
		return
	}
	Inorder(root.left)
	fmt.Print(root.val)
	fmt.Print(" ")
	Inorder(root.right)
}

// RInorder 中序遍历非递归实现
func RInorder(root *TreeNode) {
	if root == nil {
		fmt.Println("empty tree")
		return
	}
	var stack []*TreeNode
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.left
		} else {
			root = stack[len(stack)-1]
			fmt.Print(root.val)
			fmt.Print(" ")
			stack = stack[:len(stack)-1]
			root = root.right
		}
	}
}

// RPreorder 前序遍历非递归实现
func RPreorder(root *TreeNode) {
	if root == nil {
		fmt.Println("empty tree")
		return
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(root.val)
		fmt.Print(" ")
		if root.right != nil {
			stack = append(stack, root.right)
		}
		if root.left != nil {
			stack = append(stack, root.left)
		}
	}
}
