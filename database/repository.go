package database

import (
	"example/go-todo-api/structs"
)

var tasks []structs.Task = []structs.Task{}

func Insert(newTask structs.Task) {
	var task structs.Task

	task.Done = newTask.Done
	task.Name = newTask.Name
	task.Description = newTask.Description

	tasks = append(tasks, task)
}

func List() []structs.Task {
	return tasks
}

func Find(name string) (int, bool) {
	for i, task := range tasks {
		if name == task.Name {
			return i, true
		}
	}

	return -1, false
}

func Update(id int, newTask structs.Task) {
	tasks[id] = newTask
}

func Delete(id int) {
	var newArr []structs.Task = tasks[:id]

	for i := id + 1; i < len(tasks); i++ {
		newArr = append(newArr, tasks[i])
	}

	tasks = newArr
}
