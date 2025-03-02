package services

import "github.com/realjv3/gotasks/domain"

type taskService struct {
	taskRepo domain.TaskRepository
}

func NewTaskService(taskRepo domain.TaskRepository) domain.TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) CreateTask(task *domain.Task) (*domain.Task, error) {
	return s.taskRepo.Create(task)
}

func (s *taskService) GetTasksByUser(userID int) ([]*domain.Task, error) {
	return s.taskRepo.GetByUser(userID)
}

func (s *taskService) FinishTask(taskID int) (*domain.Task, error) {
	t, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return nil, err
	}

	t.Done = true

	return s.taskRepo.Update(t)
}
