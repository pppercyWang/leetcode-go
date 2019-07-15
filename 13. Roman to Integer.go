package main

import "strings"
import "fmt"

func main(){
	fmt.Println(romanToInt("VI"));
}
func romanToInt(s string)  int{
	sum :=0
	i := 0
    s = strings.ToLower(s)
    m := map[byte]int{'i':1, 'v' : 5, 'x' : 10, 'l' : 50, 'c' : 100, 'd' : 500, 'm' : 1000,};
	for i<len(s){
		if i+1<len(s) {
			if m[s[i]]<m[s[i+1]]{
				sum += m[s[i+1]] - m[s[i]]
				i += 2
				continue
			}else{
				sum += m[s[i]]
			}
		}else{
			sum += m[s[i]]
		}
		i++;
	}
    return sum
}