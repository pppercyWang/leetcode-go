package main
// import(
// 	"fmt"
// )
// func main(){
// 	isPalindrome(909)
// }

func isPalindrome(x int) bool {
    if x<0 {
        return false
    }
    if x==0{
        return true
    }
    var reverse, origin int
    origin = x
    for ;x!=0;x/=10{
        val:= x%10
        reverse = reverse*10+val
	}
    return origin==reverse
}

/*
解题思路：
只判断正整数，将其反转。若相等则返回true
*/