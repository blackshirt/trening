package trx

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type TrxRepo interface {
	Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error)
	Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error)

	Trx(ctx context.Context, obj *models.TrxDetail) (models.Trx, error)
	Organizer(ctx context.Context, obj *models.TrxDetail) (*models.Org, error)
	Location(ctx context.Context, obj *models.TrxDetail) (*models.Org, error)
	Participants(ctx context.Context, obj *models.TrxDetail) ([]models.Asn, error)

	TrxList(ctx context.Context) ([]models.TrxDetail, error)
}

type trxRepo struct {
	db *sql.DB
}

func NewTrxRepo(conn *sql.DB) TrxRepo {
	return &trxRepo{db: conn}
}

func (t *trxRepo) Category(ctx context.Context, obj *models.Trx) (*models.TrxCat, error) {
	query := "SELECT * FROM trx_category WHERE id=?"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, obj.Category.ID)
	defer stmt.Close()

	trxCat := new(models.TrxCat)
	if err := row.Scan(
		&trxCat.ID,
		&trxCat.Name,
		&trxCat.Description,
	); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return trxCat, nil
}

func (t *trxRepo) Type(ctx context.Context, obj *models.Trx) (*models.TrxType, error) {
	query := "SELECT * FROM trx_type WHERE id=?"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, obj.Type.ID)
	defer stmt.Close()

	trxType := new(models.TrxType)
	if err := row.Scan(
		&trxType.ID,
		&trxType.Name,
		&trxType.Description,
	); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return trxType, nil
}

func (t *trxRepo) Trx(ctx context.Context, obj *models.TrxDetail) (models.Trx, error) {
	query := "SELECT * FROM trx_master WHERE id=?"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
	row := stmt.QueryRowContext(ctx, obj.Trx.ID)
	defer stmt.Close()

	trx := models.Trx{}
	cat := models.TrxCat{}
	tp := models.TrxType{}
	if err := row.Scan(
		&trx.ID,
		&trx.Name,
		&trx.Description,
		&cat.ID,
		&tp.ID,
	); err != nil {
		log.Fatal(err)
		//return nil, err
	}

	trx.Category = &cat
	trx.Type = &tp
	return trx, nil
}

func (t *trxRepo) Organizer(ctx context.Context, obj *models.TrxDetail) (*models.Org, error) {
	query := "SELECT * FROM org WHERE id=?"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, obj.Organizer.ID)
	defer stmt.Close()

	trxOrg := new(models.Org)
	if err := row.Scan(
		&trxOrg.ID,
		&trxOrg.Name,
		&trxOrg.LongName,
		&trxOrg.RoadNumber,
		&trxOrg.City,
		&trxOrg.Province,
	); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return trxOrg, nil
}

func (t *trxRepo) Location(ctx context.Context, obj *models.TrxDetail) (*models.Org, error) {
	query := "SELECT * FROM org WHERE id=?"
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, obj.Location.ID)
	defer stmt.Close()

	trxLoc := new(models.Org)
	if err := row.Scan(
		&trxLoc.ID,
		&trxLoc.Name,
		&trxLoc.LongName,
		&trxLoc.RoadNumber,
		&trxLoc.City,
		&trxLoc.Province,
	); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return trxLoc, nil
}

func (t *trxRepo) Participants(ctx context.Context, obj *models.TrxDetail) ([]models.Asn, error) {
	query := `SELECT t.asn_id, asn.name, asn.nip, asn.current_job,asn.current_grade, asn.current_places
				FROM trx_asn t
				JOIN asn ON t.asn_id = asn.id
				JOIN org on asn.current_places = org.id
				WHERE t.trx_detail_id=?`

	rows, err := t.db.QueryContext(ctx, query, obj.Trx.ID)
	if err != nil {
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
			log.Fatal("error scan participants:", err)
		}
		asn.CurrentPlaces = &opd
		asns = append(asns, asn)
	}
	return asns, nil
}

func (t *trxRepo) TrxList(ctx context.Context) ([]models.TrxDetail, error) {
	query := `SELECT * FROM trx_detail`
	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	trxLists := make([]models.TrxDetail, 0)

	for rows.Next() {
		item := models.TrxDetail{}
		t := models.Trx{}
		org := models.Org{}
		loc := models.Org{}
		err = rows.Scan(
			&item.ID,
			&t.ID,
			&item.Start,
			&item.Finish,
			&org.ID,
			&loc.ID,
		)

		if err != nil {
			log.Fatal("ERROR SIR", err, item)
			return nil, err
		}
		item.Trx = t
		item.Organizer = &org
		item.Location = &loc
		trxLists = append(trxLists, item)

	}

	return trxLists, nil
}
