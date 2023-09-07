package main

import(
	"fmt";
	"strings"
)
func WordCount(s string) map[string]int {
	wrd:=strings.Fields(s)
	list:=make(map[string]int)
	for _,word:=range wrd{
		if _,ok:=list[word];ok==true{
			list[word]+=1
		} else {
			list[word]=1}
	}
	return list
}

func main(){
	fmt.Println(WordCount("I was here once"))
}