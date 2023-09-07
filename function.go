package main

import(
	"fmt"
)
//variadic parameter, take all arguments and wrap them in a slice; always last parameter om param list.
func sum(values ...int) *int{
	fmt.Println(values)
	ans:=0
	for _,v:=range values{
		ans+=v
	}
	//return ans ,just copied the result to variable s in main
	return &ans //this returns the pointer address 
}

//CAN RETURN MULTIPLE VALUES,VERY IMP
func that(a,b int) (int,int){
	sum:=a+b
	diff:=a-b
	return sum,diff
}

func main(){
	s:=sum(1,2,5,6,7)
	fmt.Println("The result is ",*s)
	sum,diff:=that(4,2)
	fmt.Printf("Sum is %v,Diff is %v \n",sum,diff)

	//Anonymous Function
	func(){
		fmt.Println("Hey I'm Anonymous Function")
	}()

	//another way
	var newfn func(float64,float64) (float64)
	newfn=func(a,b float64) (float64){
		d:=a/b
		return d
	}
	result:=newfn(4.0,6.0)
	fmt.Println(result)

	//Methods
	g:=greeter{
		greeting:"Hello",
		name:"Go",
	}
	g.greet()
}

type greeter struct{
	greeting string
	name string
}

//this a method,it is a fuction executing in a known context
//func (structname) functionname(){}; creates a copy of the struct not manipulate the same struct
//if want to manipulate struct use pointer func (g *greeter)
func (g greeter) greet(){
	fmt.Println(g.greeting,g.name)
}