package pgstore

import (
	"context"

	"github.com/NicoBernardes/taskfy.git/internal/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPgTaskStore(pool *pgxpool.Pool) PgTaskStore {
	return PgTaskStore{Queries: New(pool), Pool: pool}
}

func (pgs *PgTaskStore) CreateTask(ctx context.Context, title string, description string, priority int32) (store.Task, error) {
	task, err := pgs.Queries.CreateTask(ctx, CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, nil
}

func (pgs *PgTaskStore) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	task, err := pgs.Queries.GetTaskById(ctx, id)

	if err != nil {
		return store.Task{}, err
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time,
	}, nil
}

func (pgs *PgTaskStore) ListTasks(ctx context.Context) ([]store.Task, error) {
	tasks, err := pgs.Queries.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]store.Task, 0, len(tasks))

	for _, task := range tasks {
		result = append(result, store.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Priority:    task.Priority,
			CreatedAt:   task.CreatedAt.Time,
			UpdatedAt:   task.UpdatedAt.Time,
		})
	}

	return result, nil
}
func (pgs *PgTaskStore) UpdateTask(ctx context.Context, title string, description string, priority int32, id int32) (store.Task, error) {
	task, err := pgs.Queries.UpdateTask(ctx, UpdateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
		ID:          id,
	})
	if err != nil {
		return store.Task{}, err
	}
	return store.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt.Time,
		UpdatedAt:   task.UpdatedAt.Time, //time.now() ???
	}, nil
}
func (pgs *PgTaskStore) DeleteTask(ctx context.Context, id int32) error {
	return pgs.Queries.DeleteTask(ctx, id)
}
