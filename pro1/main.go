package main

import (
	"fmt"
	TreeNode "pro1/Tree"
)

type ListNode struct {
	val  int
	next *ListNode
}

var head *ListNode

// InsertHead 头插法
func InsertHead(nums []int) {
	if head == nil {
		head = &ListNode{val: nums[0], next: nil}
		for i := 1; i < len(nums); i++ {
			newNode := &ListNode{val: nums[i], next: head}
			head = newNode
		}
	} else {
		for i := 0; i < len(nums); i++ {
			newNode := &ListNode{val: nums[i], next: head}
			head = newNode
		}
	}
}

// PrintList 头摘法遍历链表
func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println("empty list")
		return
	}
	p := head
	for p != nil {
		fmt.Print(p.val)
		fmt.Print(" ")
		p = p.next
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	TreeNode.Insert(nums)
	TreeNode.RInorder(TreeNode.Root)
}
