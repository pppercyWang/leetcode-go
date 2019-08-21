/*
给定两个二进制字符串，返回他们的和（用二进制表示）。
输入为非空字符串且只包含数字 1 和 0。
示例 1:
输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"
*/
let addBinary = function(a, b) {
	let len = Math.max(a.length, b.length);
	let result = [];
	let carry = 0;
	let ar = a.split("").reverse();
	let br = b.split("").reverse();
	for (let i = 0; i < len; i++) {
	  let ac = ar[i] === "1" ? 1 : 0;
	  let bc = br[i] === "1" ? 1 : 0;
	  result.push((ac + bc + carry) % 2);
	  carry = Math.floor((ac + bc + carry) / 2);
	}
  
	return result
	  .concat(carry || undefined)
	  .reverse()
	  .join("");
  };