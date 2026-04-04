package user

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(databaseURL string) (UserRepository, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres connection: %w", err)
	}

	if err := pingWithRetry(db, 10, 2*time.Second); err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	repo := &PostgresUserRepository{db: db}
	if err := repo.ensureSchema(); err != nil {
		return nil, fmt.Errorf("failed to ensure schema: %w", err)
	}

	return repo, nil
}

func pingWithRetry(db *sql.DB, attempts int, delay time.Duration) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = db.Ping()
		if err == nil {
			return nil
		}
		time.Sleep(delay)
	}
	return err
}

func (r *PostgresUserRepository) ensureSchema() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE
		)
	`)
	return err
}

func (r *PostgresUserRepository) GetAll() ([]User, error) {
	rows, err := r.db.Query(`SELECT id, name, email FROM users ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	var u User
	err := r.db.QueryRow(`SELECT id, name, email FROM users WHERE id = $1`, id).Scan(&u.ID, &u.Name, &u.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *PostgresUserRepository) Create(u User) (*User, error) {
	created := User{}
	err := r.db.QueryRow(
		`INSERT INTO users(name, email) VALUES($1, $2) RETURNING id, name, email`,
		u.Name,
		u.Email,
	).Scan(&created.ID, &created.Name, &created.Email)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return &created, nil
}

func (r *PostgresUserRepository) Update(id int, u User) (*User, error) {
	updated := User{}
	err := r.db.QueryRow(
		`UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id, name, email`,
		u.Name,
		u.Email,
		id,
	).Scan(&updated.ID, &updated.Name, &updated.Email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("user not found")
	}
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate key") {
			return nil, errors.New("email already exists")
		}
		return nil, err
	}

	return &updated, nil
}

func (r *PostgresUserRepository) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user not found")
	}

	return nil
}
