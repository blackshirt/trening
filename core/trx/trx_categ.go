package trx

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)



// utility function to get single row
func (t *trxRepo) get(ctx context.Context, query string, args ...interface{}) (*models.TrxCat, error) {

	stmt, err := t.db.PrepareContext(ctx, query)
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
func (t *trxRepo) CatList(ctx context.Context) ([]*models.TrxCat, error) {
	query := "SELECT * FROM trainix_category"
	rows, err := t.db.QueryContext(ctx, query)

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
func (t *trxRepo) CatById(ctx context.Context, id int) (*models.TrxCat, error) {
	query := `SELECT * FROM trainix_category WHERE id=?`
	return t.get(ctx, query, id)
}

// get type by name
func (t *trxRepo) CatByName(ctx context.Context, name string) (*models.TrxCat, error) {
	query := `SELECT * FROM trainix_category WHERE name=?`
	return t.get(ctx, query, name)
}

// utility function to check existence of the type
func (t *trxRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trainix_category WHERE name=?`
	var catname string
	err := t.db.QueryRowContext(ctx, query, name).Scan(&catname)
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
func (t *trxRepo) CatCreate(ctx context.Context, input *models.TrxCatInput) (*models.TrxCat, error) {
	if exists := t.exists(ctx, input.Name); !exists {
		_, err := t.db.ExecContext(ctx, "INSERT INTO trainix_category(name, description) VALUES(?,?)", input.Name, input.Description)
		if err != nil {
			return nil, err
		}
	}
	row, err := t.CatByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}