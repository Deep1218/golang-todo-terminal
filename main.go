package main

func main(){
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	cmdFlags := NewCmdFlags()
	
	storage.Load(&todos)
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}