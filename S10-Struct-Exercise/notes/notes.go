package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Notes struct{
	NoteTitle string `json:"title"`
	NoteContent string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(noteTitle,noteContent string) (*Notes,error){
	if noteTitle == "" || noteContent == ""{
		return nil,errors.New("Invalid Note")
	}

	return &Notes{
		NoteTitle: noteTitle,
		NoteContent: noteContent,
		CreatedAt: time.Now(),
		},nil
}

func (n *Notes) OutputData(){
	fmt.Printf("Your note is titled %v has the following content\n\n %v",n.NoteTitle,n.NoteContent);
}

func (note Notes) SaveNotes() error{
	fileName := strings.ReplaceAll(note.NoteTitle," ","_")
	fileName = strings.ToLower(fileName) + ".json"

	json , err:= json.Marshal(note)

	if(err != nil){
		return err
	}

	return os.WriteFile(fileName, json,0644)
}