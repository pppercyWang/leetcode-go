package main

import "fmt"

func main(){
	root :=new(TreeNode)
	root.Val = 8
	root.Left = &TreeNode{Left:&TreeNode{Val:1},Right:&TreeNode{Val:4},Val:3}
	root.Right = &TreeNode{Left:&TreeNode{Val:9},Right:&TreeNode{Val:11},Val:10}
	fmt.Println(increasingBST(root))
}
type TreeNode struct {
	    Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
var ret = TreeNode{Val: 0}
var tmp = &ret
var s []int
func increasingBST(root *TreeNode) *TreeNode {
	if root.Left == nil && root.Right==nil{
		return root
	}
	inOrder(root)
	for _,v:= range s{
		tmp.Right = &TreeNode{Val:v}
		tmp = tmp.Right
	}
	return ret.Right
}
func inOrder(root *TreeNode) {
	if root == nil{
		return
	}
	inOrder(root.Left)
	s = append(s, root.Val)
	inOrder(root.Right)
}

/*
题目在本地测没问题
不知道在leetcode不能通过，而且多了很多没有的数据。有知道的朋友希望能留言给我答案
解题思路：
中序遍历，将值依次放入一个切片。然后遍历切片，设置tmp.Right等于value,tmp进位
*/