package org

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type OrgRepo interface {
	OrgById(ctx context.Context, id int) (*models.Org, error)
	OrgList(ctx context.Context) ([]*models.Org, error)
	OrgCreate(ctx context.Context, input models.OrgInput) (*models.Org, error)
}

type orgRepo struct {
	db *sql.DB
}

func NewOrgRepo(conn *sql.DB) OrgRepo {
	return &orgRepo{db: conn}
}

func (m *orgRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Org, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, args...)
	org := new(models.Org)
	err = row.Scan(
		&org.ID,
		&org.Name,
		&org.LongName,
		&org.RoadNumber,
		&org.City,
		&org.Province,
	)
	if err != nil {
		log.Fatal(err)
	}
	return org, nil
}

func (m *orgRepo) OrgById(ctx context.Context, id int) (*models.Org, error) {
	query := `SELECT * FROM org WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m *orgRepo) GetByName(ctx context.Context, name string) (*models.Org, error) {
	query := `SELECT * FROM org WHERE name=?`
	return m.getOne(ctx, query, name)
}

func (m *orgRepo) listOrg(ctx context.Context, query string, args ...interface{}) ([]*models.Org, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	orgs := make([]*models.Org, 0)
	for rows.Next() {
		org := new(models.Org)
		if err = rows.Scan(
			&org.ID,
			&org.Name,
			&org.LongName,
			&org.RoadNumber,
			&org.City,
			&org.Province,
		); err == nil {
			orgs = append(orgs, org)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return orgs, nil
}

func (m *orgRepo) OrgList(ctx context.Context) ([]*models.Org, error) {
	query := `SELECT * FROM org`
	return m.listOrg(ctx, query)
}

func (m *orgRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM org WHERE name=?`
	var orgname string
	err := m.db.QueryRowContext(ctx, query, name).Scan(&orgname)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No org with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("There is org name in db. that is = %s", orgname)
		return true
	}
}

func (m *orgRepo) OrgCreate(ctx context.Context, input models.OrgInput) (*models.Org, error) {
	exist := m.exists(ctx, input.Name)
	if !exist {
		query := `INSERT INTO org(name, long_name, road_number, city, province) VALUES(?,?,?,?,?)`
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
