package model

import (
	"database/sql"
	"fmt"
	"github.com/KoizumiHana/Exadel-Golang-2021/task5/entities"
)

type TaskModel struct {
	DB *sql.DB
}

func (TaskModel TaskModel) Create(t *entities.Task) (int64, error) {
	res, err := TaskModel.DB.Exec("insert into task (name, description, duedate, status) values (?, ?, now(), ?)", t.Name, t.Description, entities.NEW)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (TaskModel TaskModel) GetAll() (*[]entities.Task, error) {
	rows, err := TaskModel.DB.Query("select id, name, description, duedate, status from task ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]entities.Task, 0)
	for rows.Next() {
		task := entities.Task{}
		if err != rows.Scan(&task.Id, &task.Name, &task.Description, &task.DueDate, &task.Status) {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return &tasks, nil
}
func (TaskModel TaskModel) GetById(id int64) (*entities.Task, error) {
	task := entities.Task{Id: id}
	err := TaskModel.DB.QueryRow("select name, description, duedate, status from task where id = ?", id).Scan(&task.Name, &task.Description, &task.DueDate, &task.Status)
	if err != nil {
		return &entities.Task{}, err
	}
	return &task, nil
}
func (TaskModel TaskModel) Update(t *entities.Task) {
	_, err := TaskModel.DB.Exec("update task set name = ?, description = ?, duedate = ?, status = ? where id = ?", t.Name, t.Description, t.DueDate, t.Status, t.Id)
	if err != nil {
		fmt.Println(err)
	}
}

func (TaskModel TaskModel) Delete(id int64) {
	_, err := TaskModel.DB.Exec("delete from task where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
}
