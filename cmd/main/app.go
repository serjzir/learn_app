package main

import (
	"context"
	"fmt"

	"github.com/serjzir/learn_app/internal/config"
	"github.com/serjzir/learn_app/pkg/client/postgresql"
	"github.com/serjzir/learn_app/pkg/logging"
	"github.com/serjzir/learn_app/pkg/storage"
)

func main() {
	logger := logging.Init()
	logger.Info("Run learn application")
	cfg := config.GetConfig()
	logger.Info("Create new database client")
	postgreSQLClient, err := postgresql.NewClient(context.Background(), 3, *cfg)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	repository := storage.NewRepository(postgreSQLClient, logger)
	newTask := storage.Task{
		AuthorID:   0,
		AssignedID: 0,
		Title:      "Тестовая задача созданная приложением",
		Content:    "Текст тестовой задачи созданной приложением",
	}
	err = repository.Create(context.TODO(), &newTask)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Info("%v", newTask)

	allTasks, err := repository.FindAll(context.Background())
	if err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println("All tasks:")
	for _, task := range allTasks {
		fmt.Println(task)
	}

	taskByID, err := repository.FindOneById(context.Background(), 1)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println("Result search by id:")
	fmt.Println(taskByID)

	err = repository.Update(context.Background(), storage.Task{Title: "Обновление задачи2", Content: "Текст обновленной тестовой задачи2"}, 2)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	err = repository.Delete(context.Background(), 1)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	fmt.Println("Delete")
}
