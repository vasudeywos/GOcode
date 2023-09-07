package main
//Every go application structured into packages, have to declare the package its part of
import (
	"fmt"
)
//to format strings

//main application entry point
func main(){
	// var i int
	//declare variable name type
	// i=45
	// i=3
	/*Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
		Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available*/
	//j :=8.
	var n bool=true
	fmt.Printf("%v",n)
}