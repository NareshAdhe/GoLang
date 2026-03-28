package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main()  {
	title,content := getNoteData()

	userNote,err := note.New(title,content)

	if(err != nil){
		fmt.Println(err)
		return
	}
	
	userNote.Display()
	err = userNote.Save()

	if(err != nil){
		fmt.Println(err)
		return
	}
}

func getNoteData() (string,string) {
	title := getUserInput("Enter the title: ")
	
	content := getUserInput("Enter the content: ")

	return title,content
}

func getUserInput(prompt string) (string) {
	fmt.Printf("%v ",prompt)
	
	reader := bufio.NewReader(os.Stdin)

	text,err := reader.ReadString('\n')

	if(err != nil){
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text,"\n")
	text = strings.TrimSuffix(text,"\r")

	return text
}