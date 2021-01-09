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
