package trx

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)



// utility function to get single row
func (t *trxRepo) get(ctx context.Context, query string, args ...interface{}) (*models.TrxType, error) {

	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
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

// get type by id
func (t *trxRepo) TypeById(ctx context.Context, id int) (*models.TrxType, error) {
	query := `SELECT * FROM trainix_type WHERE id=?`
	return t.get(ctx, query, id)
}

// get type by name
func (t *trxRepo) TypeByName(ctx context.Context, name string) (*models.TrxType, error) {
	query := `SELECT * FROM trainix_type WHERE name=?`
	return t.get(ctx, query, name)
}

// utility function to check existence of the type
func (t *trxRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trainix_type WHERE name=?`
	var typename string
	err := t.db.QueryRowContext(ctx, query, name).Scan(&typename)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No trxType with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("Existed type name in db, that is = %s", typename)
		return true
	}
}

// type create
func (t *trxRepo) TypeCreate(ctx context.Context, input *models.TrxTypeInput) (*models.TrxType, error) {
	if exists := t.exists(ctx, input.Name); !exists {
		_, err := t.db.ExecContext(ctx, "INSERT INTO trainix_type(name, description) VALUES(?,?)", input.Name, input.Description)
		if err != nil {
			return nil, err
		}
	}
	row, err := t.TypeByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}

// list all category
func (t *trxRepo) TypeList(ctx context.Context) ([]*models.TrxType, error) {
	query := "SELECT * FROM trainix_type"
	rows, err := t.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trxTypes := make([]*models.TrxType, 0)
	for rows.Next() {
		trxType := new(models.TrxType)
		if err = rows.Scan(
			&trxType.ID,
			&trxType.Name,
			&trxType.Description,
		); err == nil {
			trxTypes = append(trxTypes, trxType)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
}
	return trxTyp}