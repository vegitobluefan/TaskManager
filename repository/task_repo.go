package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/vegitobluefan/task-manager/domain"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Save(task *domain.Task) error {
	id := uuid.New().String()
	task.ID = id
	task.Status = "pending"

	_, err := r.db.Exec(`
		INSERT INTO tasks (id, type, status, payload, result)
		VALUES ($1, $2, $3, $4, $5)
	`, task.ID, task.Type, task.Status, task.Payload, "")
	return err
}

func (r *PostgresRepo) UpdateStatus(id, status, result string) error {
	_, err := r.db.Exec(`
		UPDATE tasks
		SET status = $1, result = $2
		WHERE id = $3
	`, status, result, id)
	return err
}

func (r *PostgresRepo) GetByID(id string) (*domain.Task, error) {
	row := r.db.QueryRow(`
		SELECT id, type, status, payload, result
		FROM tasks
		WHERE id = $1
	`, id)

	task := &domain.Task{}
	err := row.Scan(&task.ID, &task.Type, &task.Status, &task.Payload, &task.Result)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *PostgresRepo) ListTasks() ([]*domain.Task, error) {
	rows, err := r.db.Query(`
		SELECT id, type, status, payload, result
		FROM tasks
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		if err := rows.Scan(&task.ID, &task.Type, &task.Status, &task.Payload, &task.Result); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
