package main
import(
	"fmt"
)

func main(){
	shortestToChar("cbbvxzbc",[]byte("c")[0])
}
var left = 10
var right =  10
func shortestToChar(S string, C byte) []int {
	s:=[]byte(S)
	l:=len(s)
	ret := []int{}
	for i:=0;i<l;i++ {
		for j:=0;j<i;j++ {
			if s[j] == C {
				left  = len(s[j:i])
			}
		}
		for k:=i;k<l;k++ {
			if s[k] == C {
				right  = len(s[i:k])
				fmt.Println(right)
			}
		}
		if left<right {
			ret = append(ret,left)
		}else{
			ret = append(ret,right)
		}
	}
	fmt.Println(ret)
	return ret
}