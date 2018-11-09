package org

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type OrgRepository interface {
	GetByID(ctx context.Context, id int) (models.Org, error)
	OrgList(ctx context.Context) ([]models.Org, error)
}

type orgRepo struct {
	db *sql.DB
}

func NewOrgRepo(conn *sql.DB) OrgRepository {
	return &orgRepo{db: conn}
}

func (m *orgRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.Org, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, args...)
	org := models.Org{}
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

func (m *orgRepo) GetByID(ctx context.Context, id int) (models.Org, error) {
	query := `SELECT * FROM org WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m orgRepo) listOrg(ctx context.Context, query string, args ...interface{}) ([]models.Org, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orgs := []models.Org{}
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
			orgs = append(orgs, *org)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return orgs, nil
}

func (m orgRepo) OrgList(ctx context.Context) ([]models.Org, error) {
	query := `SELECT * FROM org`
	return m.listOrg(ctx, query)
}
