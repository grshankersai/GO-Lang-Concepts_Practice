package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"shanker.com/interfaces/notes"
	"shanker.com/interfaces/todo"
)

// Creating an interface
// If there exisits only 1 method then the name of the interface should end with 'er' at the end.

type saver interface{
	Save() error
}

type displayer interface{
	Display() 
}

type outputtable interface{
	saver
	displayer
}

// type outputtable interface{
// 	Save() error
// 	Display()
// }


func main(){
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	noteTitle,noteContent := getNoteData()
	todoText , _ := getUserInput("Enter Todo Text: ")


	note , _ := notes.New(noteTitle,noteContent);
	todo ,err := todo.New(todoText)

	if err != nil{
		return
	}


	// note.Display()
	// todo.Display()

	// todo.Save()
	// saveData(todo)
	// saveData(note)

	outputData(note)
	outputData(todo)
}


func outputData(data outputtable) error{
	data.Display()
	return saveData(data)
}


func printSomething(value any){
	// fmt.Println(value)

	// Typing option - 1
	// switch value.(type){
	// case int:
	// 	fmt.Println("Integer: ",value)
	// case float64:
	// 	fmt.Println("Float Value: ",value)
	// case string:
	// 	fmt.Println(value)
	// }

	// Typing option 2

	typedVal,ok:=value.(int)

	if(ok){
		fmt.Println("Typed value is Integer")
		typedVal += 1
		fmt.Println("Value after increment of 1 is: ",typedVal)
	}

	
}



func saveData(data saver) error{
	err := data.Save()

	if(err != nil){
		fmt.Println("Saving issue")
	}

	fmt.Println("Saving success")
	return nil
}

func getNoteData() (string,string){
	noteTitle,_:= getUserInput("Enter the note Title")
	noteContent,_:= getUserInput("Enter the note Content")
	return noteTitle,noteContent
}
func getUserInput(prompt string) (string,error){

	
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')

	if(err != nil){
		return "" , err
	}

	text = strings.Trim(text,"\n")
	text = strings.Trim(text,"\n")

	return text,nil
	
}