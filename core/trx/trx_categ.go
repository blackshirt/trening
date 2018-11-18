package trx

import (
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type TrxCatRepo interface {
	CatById(ctx context.Context, id int) (*models.TrxCat, error)
	CatByName(ctx context.Context, name string) (*models.TrxCat, error)
}

type catRepo struct {
	db *sql.DB
}

func NewTrxCat(db *sql.DB) *TrxCatRepo {
	return &catRepo{db: db}
}

// utility function to get single row
func (t *catRepo) onetrxCat(ctx context.Contex, query string, args ...interface{}) (*models.TrxCat, error) {

	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
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

// get type by id
func (m *catRepo) CatById(ctx context.Context, id int) (*models.TrxCat, error) {
	query := `SELECT * FROM trainix_category WHERE id=?`
	return m.onetrxCat(ctx, query, id)
}

// get type by name
func (m *catRepo) CatByName(ctx context.Context, name string) (*models.TrxCat, error) {
	query := `SELECT * FROM trainix_category WHERE name=?`
	return m.onetrxCat(ctx, query, name)
}

// utility function to check existence of the type
func (m *catRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trainix_category WHERE name=?`
	var catname string
	err := m.db.QueryRowContext(ctx, query, name).Scan(&catname)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No trxCat with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("There is category name in db. that is = %s", catname)
		return true
	}
}
