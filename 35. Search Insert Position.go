package main
import(
	"fmt"
)
func main(){
	arr := []int{1,4,6}
	fmt.Println(searchInsert(arr,3));
}

func searchInsert(nums []int, target int) int {
    if len(nums) == 0 {
		return 0
	}
	insert:=0
	for i,v:= range nums{
		if v == target{
			return i
		}
		if v<target{
			insert++
		}
	}
	return insert
}