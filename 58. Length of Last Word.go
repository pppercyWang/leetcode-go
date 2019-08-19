package main
import(
	 "fmt"
)
func main(){
	fmt.Println(lengthOfLastWord("aaa  vvv"))
}
func lengthOfLastWord(s string) int {
   w, i := []int32{},0
   for _,v := range s {
	   if v == 32 {
		   w = []int32{}
	   }else{
		   w = append(w,v)
		   i = len(w)
	   }
   }
   return i
}