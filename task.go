package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
)
type Student struct{
	name string
	bits_id string
	email_id string
	branch string
}
func(std Student) prntdetails(){
	fmt.Println(std.name,std.bits_id,std.email_id,std.branch)
}

func main() {
    f,_ := excelize.OpenFile("Post_task1.xlsx")
	
	rows,_ := f.GetRows("Sheet1")

	//studentArr is just a slice of all student strucs
	var studentArr []Student

	studentInstance:=make(map[string]Student)
	//creating array from intial excel
	for i,row := range rows {

		if i==0{
			continue
		}
		
		stud:=Student{
			name:row[1],
			bits_id:row[0],
			email_id:"f"+row[0][0:5]+row[0][8:11]+"@pilani.bits-pilani.ac.in",
			branch:row[0][4:6],
		}
		studentArr = append(studentArr, stud)
		studentInstance[row[0]]=stud
	}

	//creating newfile
	newf:=excelize.NewFile()
	index,_ := newf.NewSheet("Sheet1")
	for i,studts:= range studentArr{
		j:=strconv.Itoa(i)
		newf.SetCellValue("Sheet1","A"+j,studts.bits_id)
		newf.SetCellValue("Sheet1","B"+j,studts.name)
		newf.SetCellValue("Sheet1","C"+j,studts.email_id)
		newf.SetCellValue("Sheet1","D"+j,studts.branch)
	} 
	newf.SetActiveSheet(index)
	if err := newf.SaveAs("NewPostSheet.xlsx"); err != nil {
        fmt.Println(err)
    }
	

	lastf,_ := excelize.OpenFile("Post_task1.xlsx")
	
	lrows,_ := lastf.GetRows("Sheet1")
	for i,row := range lrows {

		if i==0{
			continue
		}
		
		studer:=Student{
			name:row[1],
			bits_id:row[0],
			email_id:row[2],
			branch:row[3],
		}
		studentInstance[row[0]]=studer
	}

	//studentInstance is a map of of all student instances according to name
	studentInstance["2022B2A71111P"].prntdetails()

}
