package training

import (
	"context"
	"database/sql"

	"github.com/blackshirt/trening/models"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(url string) (Repository, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &mysqlRepository{db}, nil
}

func (r *mysqlRepository) Close() {
	r.db.Close()
}

func (r *mysqlRepository) Ping() error {
	return r.db.Ping()
}

func (r mysqlRepository) PutTraining(ctx context.Context, t models.Training) error {
	q := "INSERT INTO training(id, name, description) VALUES(?, ?, ?)"
	_, err := r.db.ExecContext(ctx, q, t.id, t.name, t.description)
	return err
}

func (r mysqlRepository) GetTraining(ctx context.Context, id string) (models.Training, error) {
	q := "SELECT id, name, description FROM training WHERE id=?"
	row := r.db.QueryRowContext(ctx, q, id)
	t := models.Training{}
	if err := row.Scan(t.id, t.name, t.description); err != nil {
		return nil, err
	}
	return t, nil
}

func (r mysqlRepository) GetTrainings(ctx context.Context, offset, limit int) ([]models.Training, error) {
	q := "SELECT id, name, description FROM training ORDER BY id DESC OFFSET ? LIMIT ?"
	rows, err := r.db.QueryContext(ctx, q, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trainings := []models.Training{}
	for rows.Next() {
		t := graph.Training{}
		if err := rows.Scan(t.id, t.name, t.description); err != nil {
			trainings := append(trainings, t)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return trainings, nil
}
