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
	//GetByID(taskID int) (*Task, error)
	//Update(task *Task) (*Task, error)
	//Delete(task *Task) error
	//GetByUser(userID int) ([]*Task, error)
}

type TaskService interface {
	CreateTask(task *Task) error
	//UpdateTask(task *Task) error
	//DeleteTask(task *Task) error
	//GetTaskByID() (*Task, error)
	//GetTasksByUser(userID int) ([]*Task, error)
}
