package storage

import (
	"database/sql"

	"github.com/realjv3/gotasks/domain"
)

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) (domain.TaskRepository, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS
    		tasks
			(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				title VARCHAR(255),
		    	description TEXT,
				done BOOLEAN DEFAULT FALSE,
		    	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			);`,
	)

	if err != nil {
		return nil, err
	}

	return &taskRepo{db: db}, nil
}

func (r *taskRepo) Create(task *domain.Task) (*domain.Task, error) {
	res, err := r.db.Exec(`INSERT INTO tasks (title, description, user_id) VALUES (?, ?, ?)`, task.Title, task.Description, task.UserID)
	if err != nil {
		return nil, err
	}

	taskID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	task.ID = int(taskID)

	return task, nil
}
