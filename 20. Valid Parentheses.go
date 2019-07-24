package main
import (
	"fmt"
)
// func main(){
// 	fmt.Println(isValid("(])"))
// }
func isValid(s string ) bool {
	array := []byte(s)
	temp := make([]byte,0)
	if len(array) == 0{
		return true
	}
	if array[0] == '}' || array[0] == ']' || array[0] == ')' {
		return false
	}
	for i:= 0;i< len(array);i++ {
		if array[i] == '{' || array[i] == '[' || array[i] == '('{
			temp = append(temp,array[i])
		}else{
			if len(temp) == 0{
				return false   // []) 避免集合越界
			}
			if array[i] == ')'{
				if temp[len(temp)-1] == '('{
					temp = temp[:len(temp)-1]
				}else{
					return false;
				}
			}else if array[i] == ']'{
				if temp[len(temp)-1] == '['{
					temp = temp[:len(temp)-1]
				}else{
					return false;
				}
			}else if array[i] == '}'{
				if temp[len(temp)-1] == '{'{
					temp = temp[:len(temp)-1]
				}else{
					return false;
				}
			}
		}
	}
	return len(temp) == 0
}
