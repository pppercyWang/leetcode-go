package main


import (
    "math"
    // "fmt"
)

// func main(){
//     fmt.Println(reverse(-234))
// }
func reverse(x int) (out int) {
    for ; x != 0; x /= 10 {
        out = out * 10 + x % 10
        if out > math.MaxInt32 || out < -math.MaxInt32 - 1 {
            return 0
        }
    }
    return out
}
/*
解题思路：
复合赋值运算符：x/=10 == x = x/10  
*/