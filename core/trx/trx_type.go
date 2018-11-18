package trx

type TrxTypeRepo interface {
	TypeById(ctx context.Context, id int) (*models.TrxType, error)
	TypeByName(ctx context.Context, name string) (*models.TrxType, error)
}

type typeRepo struct {
	db *sql.DB
}

func NewTrxType(db *sql.DB) *TrxTypeRepo {
	return &typeRepo{db: db}
}

// utility function to get single row
func (t *typeRepo) onetrxType(ctx context.Contex, query string, args ...interface{}) (*models.TrxType, error) {

	stmt, err := m.db.PrepareContext(ctx, query)
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
func (m *typeRepo) TypeById(ctx context.Context, id int) (*models.TrxType, error) {
	query := `SELECT * FROM trainix_type WHERE id=?`
	return m.onetrxType(ctx, query, id)
}

// get type by name
func (m *typeRepo) TypeByName(ctx context.Context, name string) (*models.TrxType, error) {
	query := `SELECT * FROM trainix_type WHERE name=?`
	return m.onetrxType(ctx, query, name)
}

// utility function to check existence of the type
func (m *typeRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trainix_type WHERE name=?`
	var typename string
	err := m.db.QueryRowContext(ctx, query, name).Scan(&typename)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No trxType with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("There is type name in db. that is = %s", typename)
		return true
	}
}
