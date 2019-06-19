package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	dummy.Next = nil
	tail := &dummy
	for l1 != nil ||l2 != nil{
		if l1 == nil{
			tail.Next = l2
			break
		}else if l2 == nil{
			tail.Next = l1
			break
		}else {
			if l1.Val <= l2.Val{
				tail.Next  = l1
				l1 = l1.Next
			}else {
				tail.Next = l2
				l2 = l2.Next
			}
			tail = tail.Next
		}
	}
	return dummy.Next
}
/*
解题思路：
当l1和l2不为nil时，比较l1.Val和l2.Val，谁小将它赋值给tail.Next
tail = tail.Next 保证tail永远为链表的尾部。每次将较小的数给tail。最后得到的结果则是答案
*/