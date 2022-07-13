package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/serjzir/learn_app/pkg/client/postgresql"
	"github.com/serjzir/learn_app/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	FindAll(ctx context.Context) (t []Task, err error)
	FindOneById(ctx context.Context, id int) (Task, error)
	FindOneByLabel(ctx context.Context, label string) (Task, error)
	Update(ctx context.Context, t Task, id int) error
	Delete(ctx context.Context, id int) error
}

func (r *repository) Create(ctx context.Context, t *Task) error {
	q := "INSERT INTO tasks (title, content) VALUES($1, $2) RETURNING id"
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	if err := r.client.QueryRow(ctx, q, t.Title, t.Content).Scan(&t.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQL State: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (t []Task, err error) {
	q := "SELECT id, opened, closed, author_id, assigned_id, title, content FROM tasks"
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	for rows.Next() {
		var tsk Task
		err = rows.Scan(&tsk.ID, &tsk.Opened, &tsk.Closed, &tsk.AuthorID, &tsk.AssignedID, &tsk.Title, &tsk.Content)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, tsk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *repository) FindOneById(ctx context.Context, id int) (Task, error) {
	q := "SELECT id, opened, closed, author_id, assigned_id, title, content FROM tasks WHERE id = $1"
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	var tsk Task
	err := r.client.QueryRow(ctx, q, id).Scan(&tsk.ID, &tsk.Opened, &tsk.Closed, &tsk.AuthorID, &tsk.AssignedID, &tsk.Title, &tsk.Content)
	if err != nil {
		return Task{}, err
	}
	return tsk, nil
}

func (r *repository) FindOneByLabel(ctx context.Context, label string) (Task, error) {
	q := "SELECT title, content, name FROM tasks LEFT JOIN labels ON labels.id = tasks.id where labels.name = $1;"
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	var tsk Task
	err := r.client.QueryRow(ctx, q, label).Scan(&tsk.Title, &tsk.Content)
	if err != nil {
		return Task{}, err
	}
	return tsk, nil
}

func (r *repository) Update(ctx context.Context, t Task, id int) error {
	q := "UPDATE tasks SET title=$1,content=$2 WHERE id=$3"
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Exec(ctx, q, t.Title, t.Content, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	q := "DELETE FROM tasks WHERE id = $1"
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) TaskRepository {
	return &repository{
		client: client,
		logger: logger,
	}
}
