package main

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _,v := range A{
		sum += v
	}
	if sum%3 != 0{
		return false
	}
	cur := 0
	count := 0
	for _,v := range A{
		cur += v
		if cur == sum/3{
			count++
			cur = 0
		}
	}
	return count == 3 && cur ==0
}
/*
解题思路：
1.求出总和
2.如果总和%3!=0 直接return false
*/