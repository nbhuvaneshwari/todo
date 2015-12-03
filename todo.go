package main

import (
	"flag"
	"fmt"
	"os"
)

func addItem(a1, todofilePath string) {
	todoFile, err := os.OpenFile(todofilePath, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	todoFile.Write([]byte(a1))
	todoFile.Write([]byte("\n"))
	todoFile.Close()
	fmt.Printf("+ successfully %s", a1)
}

func main() {
	todoDir, _ := os.Getwd()
	todofilePath := todoDir + "/todo.txt"
	if _, err := os.Stat(todofilePath); os.IsNotExist(err) {
		f, _ := os.Create(todofilePath)
		f.Close()
	}
	os.Setenv("TODO_FILE_PATH", todofilePath)

	addTodo := flag.String("add", "", "To add an item todo list default no item")
	flag.Parse()
	addItem(*addTodo, todofilePath)
}
