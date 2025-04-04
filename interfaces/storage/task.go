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
				title VARCHAR(255) NOT NULL,
		    	description TEXT,
				done BOOLEAN DEFAULT FALSE,
				user_id INTEGER NOT NULL,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		    	FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
			);`,
	)

	if err != nil {
		return nil, err
	}

	return &taskRepo{db: db}, nil
}

func (r *taskRepo) Create(task *domain.Task) (*domain.Task, error) {
	err := r.db.QueryRow(
		`INSERT INTO tasks (title, description, user_id) VALUES (?, ?, ?) RETURNING id`,
		task.Title,
		task.Description,
		task.UserID,
	).Scan(&task.ID)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepo) GetByID(id int) (*domain.Task, error) {
	row := r.db.QueryRow("SELECT * FROM tasks WHERE id=?", id)

	var task domain.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.UserID, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepo) GetByUser(userID int) ([]*domain.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task

	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.UserID, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *taskRepo) Update(task *domain.Task) (*domain.Task, error) {
	err := r.db.QueryRow(
		"UPDATE tasks SET title=?, description=?, done=?, updated_at = CURRENT_TIMESTAMP WHERE id=? RETURNING *",
		task.Title,
		task.Description,
		task.Done,
		task.ID,
	).Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.UserID, &task.CreatedAt, &task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return task, nil
}
