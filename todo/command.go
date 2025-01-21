package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo by giving the title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and giving new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todo")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error. Invalid index")
			os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)
	default:
		println("Invalid command")
	}
}
