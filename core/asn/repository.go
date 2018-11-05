package asn

import (
	"context"
	"database/sql"

	"github.com/blackshirt/trening/models"
)

type ASNRepository interface {
	GetByID(ctx context.Context, id int) (models.ASN, error)
  Fetch(ctx context.Context, cursor string, num int) ([]models.ASN, error)
}


type mysqlASNRepo struct {
  db *sql.DB
}

func NewMysqlASNRepo(conn *sql.DB) ASNRepository {
  return &mysqlASNRepo{db: conn}
}


func (m mysqlASNRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]models.ASN, error) {
  rows, err := m.db.QueryContext(ctx, query, args...)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  result := make([]models.ASN, 0)
  for rows.Next() {
    t := new(models.ASN)
    opdID := int(0)
    err = rows.Scan(
      &t.ID,
		   &t.Name,
		   &t.Nip,
		   &t.CurrentJob,
		   &t.CurrentGrade,
		   &opdID,
    )
    if err != nil {
      return nil, err
    }
    t.CurrentPlaces = models.OPD{
      ID: opdID,
    }
    result = append(result, t)
  }
  return result, nil
}


func (m mysqlASNRepo) Fetch(ctx context.Context, cursor string, num int) ([]models.ASN, error) {
  query := `SELECT id, name, nip, current_job, current_grade, current_places FROM asn WHERE id > ? LIMIT ?`
  return m.fetch(ctx, query, cursor, num)
}



func (m mysqlASNRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.ASN, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	a := &models.ASN{}
	opdID := int(0)
	err := row.Scan(
		&a.ID,
		&a.Name,
		&a.Nip,
		&a.CurrentJob,
		&a.CurrentGrade,
		&opdID
	)
	if err != nil {
		return nil, err
	}
 a.CurrentPlaces := models.OPD{
   ID: opdID
 }
	return a, nil
}



func (m mysqlASNRepo) GetByID(ctx context.Context, id int) (models.ASN, error) {
  query := `SELECT * FROM asn WHERE id = ?`
  return m.db.getOne(ctx, query, id)
}



