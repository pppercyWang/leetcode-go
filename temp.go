package main
import(
	"fmt"
)

func main(){
	shortestToChar("abcde",[]byte("e")[0])
}
func shortestToChar(S string, C byte) []int {
	s:=[]byte(S)
	l:=len(s)
	ret := []int{}
	left := -1
	for i:=0;i<l;i++ {
		for j:=0;j<i;j++ {
			if s[j] == s[i] {
				left  = j+1
			}
		}
	}
	fmt.Println(s)
	return ret
}