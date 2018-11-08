package orgz

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type OrgRepository interface {
	GetByID(ctx context.Context, id int) (models.Orgz, error)
	OrgList(ctx context.Context) ([]models.Orgz, error)
}

type orgRepo struct {
	db *sql.DB
}

func NewOrgRepo(conn *sql.DB) OrgRepository {
	return &orgRepo{db: conn}
}

func (m *orgRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.Orgz, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, args...)
	org := models.Orgz{}
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

func (m *orgRepo) GetByID(ctx context.Context, id int) (models.Orgz, error) {
	query := `SELECT * FROM orgz WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m orgRepo) listOrg(ctx context.Context, query string, args ...interface{}) ([]models.Orgz, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orgs := []models.Orgz{}
	for rows.Next() {
		org := new(models.Orgz)
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

func (m orgRepo) OrgList(ctx context.Context) ([]models.Orgz, error) {
	query := `SELECT * FROM orgz`
	return m.listOrg(ctx, query)
}
