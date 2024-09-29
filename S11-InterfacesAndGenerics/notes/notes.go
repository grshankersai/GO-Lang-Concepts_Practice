package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title,content string) (Note,error){
	if title == "" || content == ""{
		return Note{},errors.New("Invalid Note")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	},nil
}

func (note Note) Display(){
	fmt.Println("Note title: ",note.Title)
	fmt.Println("Note Content: ",note.Content)
}

func (note Note) Save() error{
	fileName := strings.ReplaceAll(note.Title," ","_")
	fileName = strings.ToLower(fileName)+".json"

	response,err:=json.Marshal(note)

	if err != nil{
		return err
	}

	return os.WriteFile(fileName,response,0644)

}