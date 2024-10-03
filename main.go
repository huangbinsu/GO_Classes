package main

func main() {
	api()
}

type ToDo struct {
	Task string `json:"task name"`
	Done bool   `json:"done"`
}

// 1. TODO CRUD
//-- POST create / new task 					http://127.0.0.1:800/todo/add
//-- GET read list / filter Done or Not			http://127.0.0.1:800/todo/show?status=true or http://127.0.0.1:800/todo/show?name={name}
//-- PUT update / mark task as done or not		http://127.0.0.1:800/todo/update/{name}?status=true
//-- DELETE delete / delete todo				http://127.0.0.1:800/todo/del/{name}
