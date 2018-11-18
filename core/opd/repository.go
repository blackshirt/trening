package opd

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type OpdRepo interface {
	OpdById(ctx context.Context, id int) (*models.Opd, error)
	OpdByName(ctx context.Context, name string) (*models.Opd, error)
	OpdList(ctx context.Context, limit, offset int) ([]*models.Opd, error)
	OpdCreate(ctx context.Context, input models.OpdInput) (*models.Opd, error)
}

type opdRepo struct {
	db *sql.DB
}

func NewOpdRepo(conn *sql.DB) OpdRepo {
	return &opdRepo{db: conn}
}

func (o *opdRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Opd, error) {
	stmt, err := o.db.PrepareContext(ctx, query)
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

func (o *opdRepo) OpdById(ctx context.Context, id int) (*models.Opd, error) {
	query := `SELECT * FROM opd WHERE id=?`
	return o.getOne(ctx, query, id)
}

func (o *opdRepo) OpdByName(ctx context.Context, name string) (*models.Opd, error) {
	query := `SELECT * FROM opd WHERE name=?`
	return o.getOne(ctx, query, name)
}

func (o *opdRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM opd WHERE name=?`
	var opdname string
	err := o.db.QueryRowContext(ctx, query, name).Scan(&opdname)
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

func (o *opdRepo) listOPD(ctx context.Context, query string, args ...interface{}) ([]*models.Opd, error) {
	rows, err := o.db.QueryContext(ctx, query, args...)

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

func (o *opdRepo) OpdList(ctx context.Context, limit, offset int) ([]*models.Opd, error) {
	query := `SELECT * FROM opd LIMIT ? OFFSET ?`
	return o.listOPD(ctx, query, limit, offset)
}

func (o *opdRepo) OpdCreate(ctx context.Context, input models.OpdInput) (*models.Opd, error) {
	exist := o.exists(ctx, input.Name)
	if !exist {
		query := `INSERT INTO opd(name, long_name, road_number, city, province) VALUES(?,?,?,?,?)`
		_, err := o.db.ExecContext(ctx, query, input.Name, input.LongName, input.RoadNumber, input.City, input.Province)
		if err != nil {
			return nil, err
		}
	}

	row, err := o.OpdByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}
