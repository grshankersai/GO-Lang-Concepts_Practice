package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"shanker.com/struct-exercise/notes"
)


func main(){
	noteTitle,noteContent := getNoteData()
	note , _ := notes.New(noteTitle,noteContent);

	note.OutputData()
	err  := note.SaveNotes()

	if(err != nil){
		fmt.Println("Saving note issue")
	}

	fmt.Println("Saving note success")

	// fmt.Println(note)
	// fmt.Println(noteContent)
}

func getNoteData() (string,string){
	noteTitle:= getUserInput("Enter the note Title")
	noteContent:= getUserInput("Enter the note Content")
	return noteTitle,noteContent
}

func getUserInput(prompt string) (string){
	fmt.Println(prompt)
	// var str string;
	// fmt.Scanln(&str)

	reader := bufio.NewReader(os.Stdin)
	text,err := reader.ReadString('\n')

	if err != nil{
		return ""
	}
	
	text = strings.Trim(text,"\n")
	text = strings.Trim(text,"\r")

	return text
}