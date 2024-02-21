package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	UserId    uint32
	Id        uint32
	Title     string
	Completed bool
}

func GetTask(Id int) (Task, error) {
	var task Task
	var r, err = http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", Id))
	if err != nil {
		fmt.Printf("Fetch error: %v", err)
		return task, err
	}

	//read the body
	responseData, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("IO reading error: %v", err)
		return task, err
	}

	//parse json to struct
	err = json.Unmarshal(responseData, &task)
	if err != nil {
		fmt.Printf("Error to parsing : %v", err)
		return task, err
	}

	return task, err
}
