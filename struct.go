package main

import(
	"fmt"
	"reflect"
)
type Doctor struct{
	number int
	actorName string
	companions []string
}

type Animal struct{
	Name string `required max:100`
	Origin string
}
type Bird struct{
	Animal
	SpeedKPh float32
	CanFly bool
}
func main(){
	aDoctor:=Doctor{
		number:4,
		actorName:"Jane Partake",
		companions:[]string{
			"Lid Shaw","John Grade","Sarah Jane Smith",
		},//, important to be placed at the last
	}
	fmt.Println(aDoctor)
	//second method
	aName:=struct{name string}{name:"James"}//first {  } defines structure of struct second { } initializes
	fmt.Println(aName)
	//structs not like arrays when copy structs whole data copied not just addresses

	//Go doesnt support inheritance but does support composition through embedding
	b:=Bird{}
	b.Name="Emu"
	b.Origin="Aussie"
	b.SpeedKPh=49
	b.CanFly=false
	//Bird still doesnt have any relationship to Animal;bird has animal.
	fmt.Println(b)

	//If using literal structure
	b2:=Bird{
		Animal: Animal{Name:"Kiwi",Origin: "NewZealand"},
		SpeedKPh: 23,
		CanFly: false,
	}
	fmt.Println(b2)

	//Tags using reflection;Tags only provide a string of text
	t:=reflect.TypeOf(Animal{})
	field,_:=t.FieldByName("Name")
	fmt.Println(field.Tag)
}