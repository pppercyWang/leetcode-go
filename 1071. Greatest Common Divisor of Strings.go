package main

func gcd(a int,b int)int{
	if b==0{
		return a
	}else{
		return gcd(b,a%b)
	}
}


func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1{
		return ""
	}
	s1 := []byte(str1)
	s2 := []byte(str2)
	return string(s1[0:gcd(len(s1), len(s2))])
}
/*
解题思路:
假设：str1 = ababab str2 = abab 公因子：ab
我们可以这样看成，str1等于m个ab，str2等于n个ab，所以str1+str2就是m+n个ab的重复
所以只要str1 + str2 == str2 + str1就代表该题有解

我们平时求两个数最大公约数最多用的就是辗转相除法
而str1的0索引到两个字符串长度的最大公约数即是解答
6 4 => 2
*/