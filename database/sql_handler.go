package database

import (
	"context"
	"github.com/fumist23/api-handson/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetList(ctx context.Context) (model.Tasks, error) {
	rows, err := DB.QueryContext(ctx, "SELECT Id, Title, Description  FROM Tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks model.Tasks

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			log.Printf("rows.Scan error %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask(ctx context.Context, id int) (model.Task, error) {
	row := DB.QueryRowContext(ctx, "SELECT Id, Title, Description FROM Tasks WHERE id=?", id)
	var task model.Task
	if err := row.Scan(&task.Id, &task.Title, &task.Description); err != nil {
		log.Printf("row.Scan Error: ", err)
		return task, err
	}

	return task, nil
}

func InsertTask(ctx context.Context, task model.Task) (bool, error) {
	result, err := DB.Query("INSERT INTO Tasks(Title, Description) VALUES(?, ?)", task.Title, task.Description)
}
