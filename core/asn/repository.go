package asn

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type ASNRepository interface {
	GetById(ctx context.Context, id int) (*models.Asn, error)
	AsnList(ctx context.Context) ([]*models.Asn, error)
}

type asnRepo struct {
	db *sql.DB
}

func NewASNRepo(conn *sql.DB) ASNRepository {
	return &asnRepo{db: conn}
}

func (m *asnRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Asn, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, args...)
	asn := new(models.Asn)
	err = row.Scan(
		&asn.ID,
		&asn.Name,
		&asn.Nip,
		&asn.CurrentJob,
		&asn.CurrentGrade,
		&asn.CurrentPlaces,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return asn, nil
}

func (m *asnRepo) GetById(ctx context.Context, id int) (*models.Asn, error) {
	query := `SELECT * FROM asn WHERE id = ?`
	return m.getOne(ctx, query, id)
}

func (m *asnRepo) listAsn(ctx context.Context, query string, args ...interface{}) ([]*models.Asn, error) {
	stmt, err := m.db.PrepareContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	asns := make([]*models.Asn, 0)
	for rows.Next() {
		asn := new(models.Asn)
		if err = rows.Scan(
			&asn.ID,
			&asn.Name,
			&asn.Nip,
			&asn.CurrentJob,
			&asn.CurrentGrade,
			&asn.CurrentPlaces.ID,
		); err != nil {
			log.Fatal(err)
		}
		asns = append(asns, asn)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return asns, nil
}

func (m *asnRepo) AsnList(ctx context.Context) ([]*models.Asn, error) {
	query := `SELECT * FROM asn`
	return m.listAsn(ctx, query)
}
