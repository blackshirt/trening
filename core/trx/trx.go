package trx

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type TrxRepository interface {
	TrxByID(ctx context.Context, id int) (*models.Training, error)
	TrxList(ctx context.Context, limit, offset int) ([]*models.Training, error)
	Insert(ctx context.Context, input models.Training) (*models.Training, error)
}

type trxRepo struct {
	db *sql.DB
}

func NewOPDRepo(conn *sql.DB) TrxRepository {
	return &trxRepo{db: conn}
}

func (t *trxRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Training, error) {
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	defer stmt.Close()

	trx := new(models.Training)
	if err := row.Scan(
		&trx.ID,
		&trx.Name,
		&trx.Description,
		&trx.Start,
		&trx.Finish,
		&trx.Organizer,
		&trx.Location,
		&trx.Participants,
	); err != nil {
		log.Fatal(err)
	}

	return trx, nil
}

func (t *trxRepo) TrxByID(ctx context.Context, id int) (*models.Training, error) {
	query := `SELECT * FROM trx WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (t *trxRepo) GetByName(ctx context.Context, name string) (*models.Training, error) {
	query := `SELECT * FROM trx WHERE name=?`
	return m.getOne(ctx, query, name)
}

func (t *trxRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trx WHERE name=?`
	var trxname string
	err := t.db.QueryRowContext(ctx, query, name).Scan(&trxname)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No trx with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("There is trx name in db. that is = %s", trxname)
		return true
	}
}

func (t *trxRepo) listTrx(ctx context.Context, query string, args ...interface{}) ([]*models.Training, error) {
	rows, err := t.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trxs := make([]*models.Training, 0)
	for rows.Next() {
		trx := new(models.Training)
		if err = rows.Scan(
			&trx.ID,
			&trx.Name,
			&trx.LongName,
			&trx.RoadNumber,
			&trx.City,
			&trx.Province,
		); err == nil {
			trxs = append(trxs, trx)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trxs, nil
}

func (t *trxRepo) TrxList(ctx context.Context, limit, offset int) ([]*models.Training, error) {
	query := `SELECT * FROM trx LIMIT ? OFFSET ?`
	return t.listTrx(ctx, query, limit, offset)
}

func (t *trxRepo) Insert(ctx context.Context, input models.TrainingInput) (*models.Training, error) {
	exist := t.exists(ctx, input.Name)
	if !exist {
		query := `INSERT INTO trx(name, long_name, road_number, city, province) VALUES(?,?,?,?,?)`
		_, err := m.db.ExecContext(ctx, query, input.Name, input.LongName, input.RoadNumber, input.City, input.Province)
		if err != nil {
			return nil, err
		}
	}

	row, err := t.GetByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}
