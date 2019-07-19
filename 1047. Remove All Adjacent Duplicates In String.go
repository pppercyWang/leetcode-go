package main
func removeDuplicates2(S string) string {  //跟26题函数名相同
	stack := make([]rune,0)
	for _,c:=range S{
	   if len(stack)>0 && stack[len(stack)-1] == c{
		   stack = stack[:len(stack)-1]
	   }else{
		stack = append(stack,c)
	   }
	}
	return string(stack)
  }
/*
解题思路:
利用stack的思想，将每个char放入栈顶，如果栈顶有相同的元素，就删除两个，然后继续放
*/