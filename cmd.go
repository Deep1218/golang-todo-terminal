package main

import (
	"flag"
	"strconv"
	"os"
	"fmt"
	"strings"
)

type CmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo tile")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & a new tile. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos){
	switch {
		case cf.List:
			todos.display()
		
		case cf.Add != "":
      todos.add(cf.Add)
		
		case cf.Edit != "":
			parts := strings.SplitN(cf.Edit, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Error, invalid for edit. Please use id:new_title")
				os.Exit(1)
			}
			index, err := strconv.Atoi(parts[0])
			
			if err != nil {
				fmt.Println("Error: invalid index for edit")
				os.Exit(1)
			}

			todos.edit(index, parts[1])
		
		case cf.Toggle != -1:
			todos.toggle(cf.Toggle)

		case cf.Del != -1:
			todos.delete(cf.Del)
		
		default:
			fmt.Println("Invalid command")
	}
}