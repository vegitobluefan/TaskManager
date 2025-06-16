package postgres

import (
	"database/sql"

	"github.com/vegitobluefan/task-manager/domain"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Save(task *domain.Task) error {
	_, err := r.db.Exec(`
		INSERT INTO tasks (id, type, payload, status, result)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE SET
			type = EXCLUDED.type,
			payload = EXCLUDED.payload,
			status = EXCLUDED.status,
			result = EXCLUDED.result
	`, task.ID, task.Type, task.Payload, task.Status, task.Result)
	return err
}

func (r *PostgresRepo) GetByID(id string) (*domain.Task, error) {
	row := r.db.QueryRow("SELECT id, type, payload, status, result FROM tasks WHERE id = $1", id)
	task := &domain.Task{}
	err := row.Scan(&task.ID, &task.Type, &task.Payload, &task.Status, &task.Result)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return task, err
}

func (r *PostgresRepo) List() ([]*domain.Task, error) {
	rows, err := r.db.Query("SELECT id, type, payload, status, result FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		task := &domain.Task{}
		if err := rows.Scan(&task.ID, &task.Type, &task.Payload, &task.Status, &task.Result); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *PostgresRepo) UpdateStatus(id string, status string, result string) error {
	_, err := r.db.Exec(`
		UPDATE tasks
		SET status = $1, result = $2
		WHERE id = $3
	`, status, result, id)
	return err
}
