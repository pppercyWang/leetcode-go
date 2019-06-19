package main

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	ret := []int{}
	for i,num:=range nums{
		another := target-num
		if iOfAnother,ok:=m[another];ok{
			ret =  []int{iOfAnother,i}
			break
		}
		m[num] = i
	}
	return ret
}

/*
解题思路:
每次循环将num作为键，索引作为值放入map 比如 nums[1,3,5] target=4。第一个放入的是1，0
第二次循环 target-num = 1 去判断m[1]是否存在 若存在。就找到了两个数。
*/