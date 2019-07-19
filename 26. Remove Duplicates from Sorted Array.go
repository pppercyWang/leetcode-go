package main


func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	out := 1
	for i:=1;i<len(nums);i++ {
		if nums[i-1] != nums[i] {
			nums[out] = nums[i]
			out++
		}
	}
	return out
}