package main

// import (
// 	"fmt"
// )

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	// todos.Add("Buy milk")
	// todos.Add("buy car")
	// fmt.Printf("%+v\n\n", todos)
	// todos.Print()
	// todos.Delete(0)
	// fmt.Printf("%+v\n\n", todos)
	// todos.Toggle(0)
	// todos.Print()

	storage.Save(todos)
}