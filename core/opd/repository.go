package opd

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type OPDRepository interface {
	GetByID(ctx context.Context, id int) (*models.Opd, error)
	OpdList(ctx context.Context, limit, offset int) ([]*models.Opd, error)
	Insert(ctx context.Context, input models.OPDInput) (*models.Opd, error)
}

type opdRepo struct {
	db *sql.DB
}

func NewOPDRepo(conn *sql.DB) OPDRepository {
	return &opdRepo{db: conn}
}

func (m *opdRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Opd, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, args...)
	defer stmt.Close()
	opd := new(models.Opd)
	if err := row.Scan(
		&opd.ID,
		&opd.Name,
		&opd.LongName,
		&opd.RoadNumber,
		&opd.City,
		&opd.Province,
	); err != nil {
		log.Fatal(err)
	}

	return opd, nil
}

func (m *opdRepo) GetByID(ctx context.Context, id int) (*models.Opd, error) {
	query := `SELECT * FROM opd WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m *opdRepo) GetByName(ctx context.Context, name string) (*models.Opd, error) {
	query := `SELECT * FROM opd WHERE name=?`
	return m.getOne(ctx, query, name)
}

func (m *opdRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM opd WHERE name=?`
	var opdname string
	err := m.db.QueryRowContext(ctx, query, name).Scan(&opdname)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No opd with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("There is opd name in db. that is = %s", opdname)
		return true
	}
}

func (m *opdRepo) listOPD(ctx context.Context, query string, args ...interface{}) ([]*models.Opd, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	opds := make([]*models.Opd, 0)
	for rows.Next() {
		opd := new(models.Opd)
		if err = rows.Scan(
			&opd.ID,
			&opd.Name,
			&opd.LongName,
			&opd.RoadNumber,
			&opd.City,
			&opd.Province,
		); err == nil {
			opds = append(opds, opd)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return opds, nil
}

func (m *opdRepo) OpdList(ctx context.Context, limit, offset int) ([]*models.Opd, error) {
	query := `SELECT * FROM opd LIMIT ? OFFSET ?`
	return m.listOPD(ctx, query, limit, offset)
}

func (m *opdRepo) Insert(ctx context.Context, input models.OPDInput) (*models.Opd, error) {
	exist := m.exists(ctx, input.Name)
	if !exist {
		query := `INSERT INTO opd(name, long_name, road_number, city, province) VALUES(?,?,?,?,?)`
		_, err := m.db.ExecContext(ctx, query, input.Name, input.LongName, input.RoadNumber, input.City, input.Province)
		if err != nil {
			return nil, err
		}
	}

	row, err := m.GetByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}
