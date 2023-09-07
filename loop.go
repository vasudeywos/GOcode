package main
import(
	"fmt"
)
func main(){
	//init statement executed before iterations
	sum :=0
	for i:=0;i<10;i++{
		sum += i
	}
	fmt.Printf("%v \n",sum)
	a:=5//0101
	b:=6//0110
	//bitwise operators
	n:=13.7e73//Exponential
	s:="This is a string"
	s+=" and this is added"//take the original memory new chunk memory created,added string,and s goona get changed to point to new block of memory
	byt:=[]byte(s)//Convert to ASCII bytes
	fmt.Println(a & b)//0100
	// fmt.Println(a | b)//0111
	// fmt.Println(a ^ b)
	// fmt.Println(a &^ b)
	fmt.Println(n)
	fmt.Println(s,byt)
	fmt.Printf("%v",s[:6])
}