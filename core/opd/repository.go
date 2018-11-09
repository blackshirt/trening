package opd

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type OPDRepository interface {
	GetByID(ctx context.Context, id int) (models.OPD, error)
	OPDList(ctx context.Context) ([]models.OPD, error)
	Insert(ctx context.Context, input models.OPDInput) (int, error)
}

type opdRepo struct {
	db *sql.DB
}

func NewOPDRepo(conn *sql.DB) OPDRepository {
	return &opdRepo{db: conn}
}

func (m *opdRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.OPD, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, args...)
	defer stmt.Close()
	opd := models.OPD{}
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


func (m *opdRepo) GetByID(ctx context.Context, id int) (models.OPD, error) {
	query := `SELECT * FROM opd WHERE id=?`
	return m.getOne(ctx, query, id)
}


func (m *opdRepo) exists(ctx context.Context, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM opd WHERE name=?`
	err := m.db.QueryRowContext(ctx, query, name)
	if err != nil {
		return false, err
	}
	return true, nil
}


func (m *opdRepo) listOPD(ctx context.Context, query string, args ...interface{}) ([]models.OPD, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	opds := make([]models.OPD, 0)
	for rows.Next() {
		opd := models.OPD{}
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


func (m *opdRepo) OPDList(ctx context.Context) ([]models.OPD, error) {
	query := `SELECT * FROM opd`
	return m.listOPD(ctx, query)
}


func (m *opdRepo) Insert(ctx context.Context, input models.OPDInput) (int, error) {
 exist, err := m.exists(ctx, input.Name)
 if !exist {
   query := `INSERT INTO opd(name, long_name, road_number, city, province) VALUES(?,?,?,?,?)`
	  res, err := m.db.ExecContext(ctx, query, input.Name, input.LongName, input.RoadNumber, input.City, input.Province)
	  if err != nil {
	    return nil, err
	  }
	  return int(res.LastInsertId()), nil
	}
	return nil, err
}
