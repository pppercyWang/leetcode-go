package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int) 
	start := 0  
	maxLength := 0  
	for i,char := range []byte(s){
		if charI,ok := m[char]; ok && charI>=start{  
			start = m[char]+1
		}
		if i- start +1 > maxLength{ 
			maxLength = i - start +1
		}
		m[char] = i
	}
	return maxLength
}
/*
变量：
m :储存每个字符在这个字符串最后一次出现的索引。键为char，值为索引
start :每次可能是最长子字符串开始位置的索引
maxLength :最大长度
if_1：如果这个字母在m中存在，且m[char]>=start时，start等于m[char]+1。当m[char]<start我们

解题思路：
将字符串的每一个字符依次放入map中，键为char 值为出现在这个字符串的索引。如果遇到重复的字符，就会覆盖掉map中之前字符的值。所以保证了map中存放的都是每个字符的最后出现在字符串的索引。
因为最长字串是连续的。所以这里如果遇到重复的字符。start就需要改变
这里有三种情况:
看最后一次循环
第一种start=m[char](m[char]为上一个重复字符 例如abca m[char]就为0),此时start需要进一位。 最长子串变成bca
第二种start<mchar,此时需要进一位。最长子串为cb
第三种start>mchar,m[char]为0，但start为2,这时第一个a已经与最长字串没有关系了。所以start不需要改变。最长字串为cba

例子：
abc
一直不会执行if_1 每次maxLen++
abcac
最后一次循环会将start变成1 但maxLen不会变
abcb
第4次循环会将start变成2 但maxLen不会变 所以说start为 可能是最长子字符串开始位置的索引
最长字串：abc
abcbdef
第5次循环 不会走if_1 ，maxLen === i- start +1，maxLen不变 从第六次循环 maxLen++
最长字串：cbdef
*/