package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("in.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	fmt.Println("File:\n", string(data))
	fmt.Println("Number of bytes read:", len(data))

	fileInfo, err := os.Stat("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

	fileCreate, errCreate := os.Create("out.txt")
	if errCreate != nil {
		fmt.Println("Error creating file", err)
		return
	}
	fileCreate.WriteString("Europe")
	fileCreate.Close()

	fileAppend, errAppend := os.OpenFile("out.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if errAppend != nil {
		fmt.Println("Error appending file", err)
		return
	}
	fileAppend.WriteString("\nAmerica")
	fileAppend.Close()

	//Get the context of a directory
	dir, errDir := os.Open(".")
	if errDir != nil {
		fmt.Println("Error reading directory", err)
		return
	}
	defer dir.Close()
	filesInfo, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	fmt.Println("Current Directory:")
	for _, fileValue := range filesInfo {
		fmt.Println(fileValue.Name())
	}
}
