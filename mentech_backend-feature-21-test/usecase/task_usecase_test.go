package usecase

import (
	"mentech-backend/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// モックの定義
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	args := m.Called(tasks, userId)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	args := m.Called(task, userId, taskId)
	return args.Error(0)
}

func (m *MockTaskRepository) CreateTask(task *model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	args := m.Called(task, userId, taskId)
	return args.Error(0)
}

func (m *MockTaskRepository) DeleteTask(userId uint, taskId uint) error {
	args := m.Called(userId, taskId)
	return args.Error(0)
}

type MockTaskValidator struct {
	mock.Mock
}

func (m *MockTaskValidator) TaskValidate(task model.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

// テスト: GetAllTasks
func TestGetAllTasks(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockValidator := new(MockTaskValidator)

	mockTasks := []model.Task{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
	}

	// モックの動作を定義
	mockRepo.On("GetAllTasks", mock.Anything, uint(1)).Return(nil).Run(func(args mock.Arguments) {
		*args.Get(0).(*[]model.Task) = mockTasks
	})

	taskUsecase := NewTaskUsecase(mockRepo, mockValidator)
	tasks, err := taskUsecase.GetAllTasks(1)

	// アサーション
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "Task 1", tasks[0].Name)
	mockRepo.AssertExpectations(t)
}

// テスト: CreateTask
func TestCreateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockValidator := new(MockTaskValidator)

	newTask := model.Task{
		ID:            1,
		Name:          "New Task",
		PurchasePrice: 100,
		Quantity:      2,
	}

	mockValidator.On("TaskValidate", newTask).Return(nil)
	mockRepo.On("CreateTask", &newTask).Return(nil)

	taskUsecase := NewTaskUsecase(mockRepo, mockValidator)
	createdTask, err := taskUsecase.CreateTask(newTask)

	assert.NoError(t, err)
	assert.Equal(t, "New Task", createdTask.Name)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

// テスト: UpdateTask
func TestUpdateTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockValidator := new(MockTaskValidator)

	existingTask := model.Task{
		ID:            1,
		Name:          "Existing Task",
		PurchasePrice: 100,
		Quantity:      2,
	}

	mockValidator.On("TaskValidate", existingTask).Return(nil)
	mockRepo.On("UpdateTask", &existingTask, uint(1), uint(1)).Return(nil)

	taskUsecase := NewTaskUsecase(mockRepo, mockValidator)
	updatedTask, err := taskUsecase.UpdateTask(existingTask, 1, 1)

	assert.NoError(t, err)
	assert.Equal(t, "Existing Task", updatedTask.Name)
	mockRepo.AssertExpectations(t)
	mockValidator.AssertExpectations(t)
}

// テスト: DeleteTask
func TestDeleteTask(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockValidator := new(MockTaskValidator)

	mockRepo.On("DeleteTask", uint(1), uint(1)).Return(nil)

	taskUsecase := NewTaskUsecase(mockRepo, mockValidator)
	err := taskUsecase.DeleteTask(1, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
