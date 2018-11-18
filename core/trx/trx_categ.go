package trx

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type CatRepo interface {
	CatById(ctx context.Context, id int) (*models.TrxCat, error)
	CatByName(ctx context.Context, name string) (*models.TrxCat, error)
	CatList(ctx context.Context) ([]*models.TrxCat, error)
	CatCreate(ctx context.Context, input *models.TrxCatInput) (*models.TrxCat, error)
}

type catRepo struct {
	db *sql.DB
}

func NewTrxCat(db *sql.DB) CatRepo {
	return &catRepo{db: db}
}

// utility function to get single row
func (c *catRepo) onetrxCat(ctx context.Context, query string, args ...interface{}) (*models.TrxCat, error) {

	stmt, err := c.db.PrepareContext(ctx, query)
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

// list all category
func (c *catRepo) CatList(ctx context.Context) ([]*models.TrxCat, error) {
	query := "SELECT * FROM trainix_category"
	rows, err := c.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trxCats := make([]*models.TrxCat, 0)
	for rows.Next() {
		trxCat := new(models.TrxCat)
		if err = rows.Scan(
			&trxCat.ID,
			&trxCat.Name,
			&trxCat.Description,
		); err == nil {
			trxCats = append(trxCats, trxCat)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trxCats, nil
}

// get category by id
func (c *catRepo) CatById(ctx context.Context, id int) (*models.TrxCat, error) {
	query := `SELECT * FROM trainix_category WHERE id=?`
	return c.onetrxCat(ctx, query, id)
}

// get type by name
func (c *catRepo) CatByName(ctx context.Context, name string) (*models.TrxCat, error) {
	query := `SELECT * FROM trainix_category WHERE name=?`
	return c.onetrxCat(ctx, query, name)
}

// utility function to check existence of the type
func (c *catRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trainix_category WHERE name=?`
	var catname string
	err := c.db.QueryRowContext(ctx, query, name).Scan(&catname)
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

// type create
func (c *catRepo) CatCreate(ctx context.Context, input *models.TrxCatInput) (*models.TrxCat, error) {
	if exists := c.exists(ctx, input.Name); !exists {
		_, err := c.db.ExecContext(ctx, "INSERT INTO trainix_category(name, description) VALUES(?,?)", input.Name, input.Description)
		if err != nil {
			return nil, err
		}
	}
	row, err := c.CatByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}
