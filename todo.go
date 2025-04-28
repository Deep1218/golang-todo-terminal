package main

import (
	"fmt"
	"time"
	"errors"
	"strconv"
	"os"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
	UpdatedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string){
	todo := Todo{
		Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}

	*todos = append(*todos, todo)
}


func (todos *Todos) validateIndex(idx int) error {
	if idx < 0 || idx >= len(*todos){
		err := errors.New("Invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) delete(idx int) error {
	t:=*todos

	if err := t.validateIndex(idx); err !=nil{
		return err
	}

	*todos = append(t[:idx], t[idx+1:]...)

	return nil
}

func (todos *Todos) toggle(idx int) error {
	t:=*todos

	if err := t.validateIndex(idx); err !=nil{
		return err
	}

	isCompleted := t[idx].Completed

	if !isCompleted{
		completionTime := time.Now()
		t[idx].CompletedAt = &completionTime
		t[idx].UpdatedAt = &completionTime
	}

	t[idx].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(idx int, title string) error {
	t:=*todos

	if err := t.validateIndex(idx); err !=nil{
		return err
	}

	currentTime := time.Now()
	t[idx].Title = title
	t[idx].UpdatedAt = &currentTime 
	return nil
}

func (todos *Todos) display(){
	table:=table.New(os.Stdout)
	// table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed","Created At", "Updated At", "Completed At")

	for index, t := range *todos{
		completed := "No"
		completedAt := ""
		updatedAt := ""

		if t.Completed {
			completed = "Yes"
			if t.CompletedAt != nil{
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		if t.UpdatedAt != nil{
			updatedAt = t.UpdatedAt.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), updatedAt, completedAt)
	}

	table.Render()
}