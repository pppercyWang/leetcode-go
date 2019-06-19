package main
type ListNode struct {
	   Val int
	    Next *ListNode
	}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode) //{0,<nil>}
	cur := ret 
	add := 0
	for l1!=nil || l2!=nil || add !=0{
		val1,val2 := 0,0
		if l1!=nil{
			val1,l1 = l1.Val,l1.Next
		}
		if l2!=nil{
			val2,l2 = l2.Val,l2.Next
		}
		sum := val1+val2+add
		add = sum/10   //判断是否需要进位
		cur.Next = &ListNode{Val:sum%10,Next:nil} //第一次循环cur指向ret。所以为ret.Next
												  //总和取余则为新节点的Val
		cur = cur.Next //保证cur为ret的最后一个Next为nil节点
	}
	return ret.Next
}
