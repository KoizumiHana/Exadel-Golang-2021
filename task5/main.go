package main

import (
	"encoding/json"
	"fmt"
	"github.com/KoizumiHana/Exadel-Golang-2021/task5/config"
	"github.com/KoizumiHana/Exadel-Golang-2021/task5/entities"
	"github.com/KoizumiHana/Exadel-Golang-2021/task5/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/task/", taskHandlerWithPathVariable)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func taskHandler(writer http.ResponseWriter, request *http.Request) {
	db, err := config.GetMysqlDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	taskModel := model.TaskModel{
		DB: db,
	}
	writer.Header().Set("Content-Type", "application/json")
	switch request.Method {
	case http.MethodPost:
		reqBody, _ := ioutil.ReadAll(request.Body)
		var task entities.Task
		json.Unmarshal(reqBody, &task)
		_, err := taskModel.Create(&task)
		if err != nil {
			return
		}
		writer.WriteHeader(http.StatusCreated)
	case http.MethodPut:
		reqBody, _ := ioutil.ReadAll(request.Body)
		var newTask entities.Task
		json.Unmarshal(reqBody, &newTask)
		taskModel.Update(&newTask)
		writer.WriteHeader(http.StatusOK)
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

func taskHandlerWithPathVariable(writer http.ResponseWriter, request *http.Request) {
	db, err := config.GetMysqlDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	taskModel := model.TaskModel{
		DB: db,
	}
	writer.Header().Set("Content-Type", "application/json")
	pathVariable := strings.TrimPrefix(request.URL.Path, "/task/")
	switch request.Method {
	case http.MethodGet:
		if pathVariable == "all" {
			tasks, err := taskModel.GetAll()
			if err != nil {
				return
			}
			json.NewEncoder(writer).Encode(tasks)
		} else if i, err := strconv.ParseInt(pathVariable, 10, 64); err == nil {
			task, err := taskModel.GetById(i)
			if err != nil {
				return
			}
			json.NewEncoder(writer).Encode(task)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		if i, err := strconv.ParseInt(pathVariable, 10, 64); err == nil {
			taskModel.Delete(i)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}
