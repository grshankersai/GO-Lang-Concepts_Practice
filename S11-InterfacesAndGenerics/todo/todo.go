package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct{
	Content string `json:"text"`
}

func New(content string) (Todo,error){
	if content == ""{
		return Todo{},errors.New("Empty content")
	}

	return Todo{
		Content: content,
	},nil
}

func (todo Todo) Display(){
	fmt.Print(todo)
}

func (todo Todo) Save() error{
	fileName := "todo.json"

	json,err:=json.Marshal(todo)

	if(err != nil){
		return err
	}

	return os.WriteFile(fileName,json,0644) 

}
