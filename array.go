package main

import(
	"fmt"
)
func main(){
	grades:=[3]int{56,72,92}//array can hold only one data type
	newarr:=[...]int{34,67}//... fill in as much arguments in array
	fmt.Printf("%v %v",grades,newarr)

	a:=make([]int,2)//make(type,len)
	fmt.Printf("Capacity:%v",cap(a))
	a= append(a,1)//append(sorce slice,element)
	fmt.Printf("Capacity:%v",cap(a))// if there is size then mutated elsegets copied and doubled

	newa:=[...]int{1,2,3,4,5}
	cp:=newa[:len(newa)-1]
	fmt.Printf("%v",cp)
	d:=append(newa[:2],newa[3:]...)//... spread out
    fmt.Printf("%v",d)
	//a gets changed though

	//Maps
	statepop:=make(map[string]int)
	statepop=map[string]int{
		"California":5647383,
		"Texas":543729,
		"Ohio":4674844,
	}
	_,ok:=statepop["Washington"]
	//_ write only variable variable to throw away the value
	//only one type of key any type for value
	fmt.Println(statepop["Texas"])
	fmt.Println(ok)
}