package domain

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	UserID      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TaskRepository interface {
	Create(task *Task) (*Task, error)
	GetByID(id int) (*Task, error)
	GetByUser(userID int) ([]*Task, error)
	Update(task *Task) (*Task, error)
}

type TaskService interface {
	CreateTask(task *Task) (*Task, error)
	GetTasksByUser(userID int) ([]*Task, error)
	FinishTask(taskID int) (*Task, error)
}
