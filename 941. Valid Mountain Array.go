package main
func validMountainArray(A []int) bool {
	if len(A)<3{
		return false
	}
	for i:=1;i<len(A)-1;i++{
		if A[i]<=A[i-1]{
			return false
		}else if A[i+1]>A[i]{  //从左往右直到找到峰值
			continue
		}
		j := i+1
		if A[j] == A[j-1]{ //如果峰值的下一位等于峰值，返回false
			return false
		}
		for ;j<len(A)-1;j++{
			if A[j]<A[j-1] && A[j+1]<A[j]{ 
				continue
			}
			return false
		}
		if j == len(A)-1 {
			return true
		}
	}
	return false
}
/*
解题思路：
先从左往右直到找到峰值，并保证峰值左边的数都是升序的，峰值右侧的都是降序的。峰值的左边和右边都不能有相同的两个数。如果不满足这三点 return false
*/