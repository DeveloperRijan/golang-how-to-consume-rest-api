package main

import (
	restapi "exapp/rest-api"
	"fmt"
)

func main() {
	var task, err = restapi.GetTask(1)

	if err != nil {
		fmt.Printf("The to fetch task: %v", err)
		return
	}

	fmt.Printf("Task Detail\nTitle: %v,\nId: %v,\nUserId: %v", task.Title, task.Id, task.UserId)
}
