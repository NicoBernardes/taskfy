package services

import (
	"testing"
	"time"

	"github.com/NicoBernardes/taskfy.git/internal/store"
	"github.com/go-jose/go-jose/v4/testutils/assert"
)

type MockTaskStore struct {
}

func (m *MockTaskStore) CreateTask(title, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) GetTaskById(id int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       "Mock Test Task",
		Description: "Mock Test Task",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) ListTasks() ([]store.Task, error) {
	return []store.Task{
		{Id: 1,
			Title:       "Mock Test Task",
			Description: "Mock Test Task",
			Priority:    1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Id:          2,
			Title:       "Mock Test Task 2",
			Description: "Mock Test Task 2",
			Priority:    -1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}, nil
}

func (m *MockTaskStore) UpdateTask(id int32, title string, description string, priority int32) (store.Task, error) {
	return store.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (m *MockTaskStore) DeleteTask(id int32) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	//Arrange
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)
	//Act
	task, err := taskService.Store.CreateTask("Mock Test Task", "Mock Test Description", 1)
	//Assert
	assert.NoError(t, err)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Test Description", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestGetTask(t *testing.T) {
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	task, err := taskService.Store.GetTaskById(1)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), task.Id)
	assert.Equal(t, "Mock Test Task", task.Title)
}

func TestListTask(t *testing.T) {
	mockStore := MockTaskStore{}
	taskService := NewTaskService(&mockStore)

	tasks, err := taskService.Store.ListTasks()

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "Mock Test Task", tasks[0].Title)
	assert.Equal(t, "Mock Test Task 2", tasks[1].Title)
}
