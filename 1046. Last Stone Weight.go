package main
import "sort"
func lastStoneWeight(stones []int) int {
	if len(stones)<2{
	   return stones[0]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(stones))) //降序排序
	not_mashed := []int{}
	for i:=0;i<len(stones)-1;i++{
	   if(i == 0){
		  not_mashed = append(not_mashed,stones[0]-stones[1])
		  not_mashed = append(not_mashed,stones[2:]...)
	   }else{
		  sort.Sort(sort.Reverse(sort.IntSlice(not_mashed)))
		  not_mashed1 := []int{not_mashed[0]-not_mashed[1]}
		  not_mashed = append(not_mashed1,not_mashed[2:]...)
	   }
 
	}
	return not_mashed[0]
 }
