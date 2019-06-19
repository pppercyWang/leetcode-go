package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1,nums2...)
	s = quickSort(s)
	lens :=len(s)
	if lens%2 == 0{
		return (float64(s[lens/2-1])+float64(s[lens/2]))/2
	}else{
		return float64(s[lens/2])
	}
}
func quickSort(arr []int) []int{
	if len(arr)<=1{
		return arr
	}
	p :=arr[0:1][0]
	temp :=arr[1:]
	left := []int{}
	right := []int{}
	for _,v :=range temp{
		if v<p{
			left = append(left, v)
		}else{
			right = append(right,v)
		}
	}
	return append(append(quickSort(left), p), quickSort(right)...)
}
/*
解题思路:
先将两个数组合并，再通过快速排序，最后找出中位数 但是我的这个解法时间复杂度不符合要求。但是比较容易理解。这个题的考的就是排序的速度
*/