package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)
type student struct{
	Name string
	BITS_ID string
	Email_Id string
	Branch string
}
func (std student) getstd(){
	fmt.Println(std.Name,std.BITS_ID,std.Email_Id,std.Branch)
}
func (std student) prntdetails(){
	fmt.Printf("Name:%v,BITS ID:%v,Email:%v,Branch:%v",std.Name,std.BITS_ID,std.Email_Id,std.Branch)
}
func main() {
    f,_ := excelize.OpenFile("Post_task1.xlsx")
	
	rows,_ := f.GetRows("Sheet1")
	
	var studentArr []student

	studentInstance:=make(map[string]student)
	
	row1:=true
	for _,row := range rows {

		if row1==true{
			row1=false
			continue
		}
		
		stud:=student{
			Name:row[1],
			BITS_ID:row[0],
			Email_Id:"f"+row[0][0:5]+row[0][8:11]+"@pilani.bits-pilani.ac.in",
			Branch:row[0][4:6],
		}
		studentArr = append(studentArr, stud)
		studentInstance[row[1]]=stud
	}

	//studentArr is just a slice of all student strucs
	for _,stdrw:=range studentArr{
	 	stdrw.getstd()
	 }

	//studentInstance is a map of of all student instances according to name
	studentInstance["AARYAMAN"].prntdetails()

}
