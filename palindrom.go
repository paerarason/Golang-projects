package main
import (
	//"fmt"
    "strings"
)

func numPalindrome(n int) bool {
	var num int=0 
	for i:=n;i>0;i/=10{
		res:=i%10
		num=10*num+res
	}
	if num==n{
		return true
	}else{
		return false 
	}
}

func StringPalindrome(s string) bool {
	 word:=strings.ToLower(strings.ReplaceAll(s, " ", ""))
	 start:=0
	 end:=len(s)
	 for start>end{
		if word[start] != word[end] {
			return false
		} 
		start++
		end--
	 }
	 return true
}








