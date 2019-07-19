package main

import "fmt"
import "strings"

// func main(){
// 	strs := []string{"aa","asc"}; 
// 	fmt.Println(longestCommonPrefix(strs))
// }
func longestCommonPrefix(strs []string) string {
	result := strings.Builder{}
	if len(strs) == 0{
		return result.String()
	}
	for i:=0;i<len(strs[0]);i++{
		c := strs[0][i]
		same :=true;
		for _,s := range strs{
			if i>=len(s) || s[i] !=c {
				same = false;
				break;
			}
		}
		if !same{
			break;
		}
		result.WriteByte(c)
	}
	return result.String()
}
/*
解题思路：
随意取字符串数组中的一个字符串，如果它是长度最短的，能找到公共前缀，最长的也能。所以这里随意找就好了
如果当前字符串长度比第一个字符串短，或者当前的s[i] !=c，则不相等，不向builder添加字符
最后return
*/