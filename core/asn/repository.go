package asn

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type AsnRepo interface {
	AsnById(ctx context.Context, id int) (*models.Asn, error)
	AsnByNip(ctx context.Context, nip string) (*models.Asn, error)
	AsnList(ctx context.Context) ([]models.Asn, error)
}

type asnRepo struct {
	db *sql.DB
}

func NewAsnRepo(conn *sql.DB) AsnRepo {
	return &asnRepo{db: conn}
}

func (a *asnRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Asn, error) {
	stmt, err := a.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal("Error get one", err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, args...)

	asn := new(models.Asn)
	err = row.Scan(
		&asn.ID,
		&asn.Name,
		&asn.Nip,
		&asn.CurrentJob,
		&asn.CurrentGrade,
		&asn.CurrentPlaces.ID,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return asn, nil
}

func (a *asnRepo) AsnById(ctx context.Context, id int) (*models.Asn, error) {
	query := `SELECT * FROM asn WHERE id = ?`
	return a.getOne(ctx, query, id)
}

func (a *asnRepo) AsnByNip(ctx context.Context, nip string) (*models.Asn, error) {
	query := `SELECT * FROM asn WHERE nip = ?`
	return a.getOne(ctx, query, nip)
}

func (a *asnRepo) listAsn(ctx context.Context, query string, args ...interface{}) ([]models.Asn, error) {
	stmt, err := a.db.PrepareContext(ctx, query)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	asns := make([]models.Asn, 0)
	for rows.Next() {
		asn := models.Asn{}
		opd := models.Opd{}
		if err = rows.Scan(
			&asn.ID,
			&asn.Name,
			&asn.Nip,
			&asn.CurrentJob,
			&asn.CurrentGrade,
			&opd.ID,
		); err != nil {
			log.Fatal(err)
		}
		asn.CurrentPlaces = &opd
		asns = append(asns, asn)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return asns, nil
}

func (a *asnRepo) AsnList(ctx context.Context) ([]models.Asn, error) {
	query := `SELECT * FROM asn`
	return a.listAsn(ctx, query)
}
