package ListNode

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

var head *ListNode

// InsertBack 尾插法
func InsertBack(nums []int) {
	if head == nil {
		head = &ListNode{val: nums[0], next: nil}
		p := head
		for i := 1; i < len(nums); i++ {
			p.next = &ListNode{val: nums[i], next: nil}
			p = p.next
		}
	} else {
		p := head
		for p.next != nil {
			p = p.next
		}
		for i := 0; i < len(nums); i++ {
			p.next = &ListNode{val: nums[i], next: nil}
			p = p.next
		}
	}
}

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

// DeleteNode 删除指定值节点
func DeleteNode(head *ListNode, val int) {
	if head == nil {
		return
	}
	if head.val == val {
		head = head.next
		return
	}
	p := head
	for p.next != nil {
		if p.next.val == val {
			p.next = p.next.next
			return
		}
		p = p.next
	}
}

// DeleteIndexNode 删除指定索引节点
func DeleteIndexNode(head *ListNode, index int) {
	if head == nil {
		return
	}
	if index == 0 {
		head = head.next
		return
	}
	p := head
	for i := 0; i < index; i++ {
		p = p.next
	}
	p.next = p.next.next
}
