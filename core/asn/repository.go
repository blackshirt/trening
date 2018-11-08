package asn

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type ASNRepository interface {
	GetByID(ctx context.Context, id int) (models.ASN, error)
	ASNList(ctx context.Context) ([]models.ASN, error)
}

type asnRepo struct {
	db *sql.DB
}

func NewASNRepo(conn *sql.DB) ASNRepository {
	return &asnRepo{db: conn}
}

func (m *asnRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.ASN, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, args...)
	asn := models.ASN{}
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
	}
	return asn, nil
}

func (m *asnRepo) GetByID(ctx context.Context, id int) (models.ASN, error) {
	query := `SELECT * FROM asn WHERE id = ?`
	return m.getOne(ctx, query, id)
}

func (m *asnRepo) listASN(ctx context.Context, query string, args ...interface{}) ([]models.ASN, error) {
	stmt, err := m.db.PrepareContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	asns := []models.ASN{}
	for rows.Next() {
		asn := models.ASN{}
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

func (m *asnRepo) ASNList(ctx context.Context) ([]models.ASN, error) {
	query := `SELECT * FROM asn`
	return m.listASN(ctx, query)
}
