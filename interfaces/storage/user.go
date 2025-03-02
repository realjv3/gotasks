package storage

import (
	"database/sql"
	"errors"

	"github.com/realjv3/gotasks/domain"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) (domain.UserRepository, error) {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS
    		users
			(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				name VARCHAR(255),
		    	email VARCHAR(255) UNIQUE NOT NULL,
		    	password VARCHAR(255) UNIQUE NOT NULL,
				active BOOLEAN DEFAULT TRUE NOT NULL,
				created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
				updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
			)`,
	)

	if err != nil {
		return nil, err
	}

	return &userRepo{db: db}, nil
}

func (r *userRepo) Create(user *domain.User) (*domain.User, error) {
	res, err := r.db.Exec(`INSERT INTO users (name, email, password) VALUES (?, ?, ?)`, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.Get(int(userID))
}

func (r *userRepo) Get(id int) (*domain.User, error) {
	row := r.db.QueryRow("SELECT * FROM users WHERE id = ?", id)

	var ret domain.User
	err := row.Scan(&ret.ID, &ret.Name, &ret.Email, &ret.Password, &ret.Active, &ret.CreatedAt, &ret.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &ret, nil
}
