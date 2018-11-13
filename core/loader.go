package loader

import (

	"context"
	"database/sql"
	"sync"
)


type TrxLoader struct {
	data map[in]*Trx
	loaded bool
	trxRepo *trx.TrxRepository
	mutex sync.Mutex
}


func NewTrxLoader(trxRepo *trx.TrxRepository) *TrxLoader {
	return &TrxLoader{
		data: make(map[int]*Trx),
		trxRepo: trxRepo,
		mutex: sync.Mutex{},
	}

}


func (t *TrxLoader) Enqueue(trxIDs []int) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	for _, trxID := range trxIDs {
		t.data[trxID] = &Trx{}
	}
}


func (t *TrxLoader) Query(ctx context.Context, trxID int) (*Trx, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	if t.loaded {
		return t.data[trxID], nil
	}

	if len(t.data) == 0 {
		return nil, nil
	}

	var trxIDs []int
	for trxID := range t.data {
		trxIDs = append(trxIDs, trxID)
	}
	err := t.load(trxIDs)
	if err != nil {
		return nil, err
	}

	t.loaded = true

	asnTrxLoader := ctx.Value(asnTrxLoaderKey).(*AsnTrxLoader)
	if asnTrxLoader != nil {
		asnTrxLoader.Enqueue(trxIDs)
	}

	return t.data[trxID], nil
}


func (t *TrxLoader) load(trxIDs []int) error {
	rows, err := t.db.Query(
		"SELECT id, name FROM trx WHERE id = ANY($1::CHAR(27)[])",
		pq.Array(trxIDs),
	)
	if err != nil {
		return err
	}

	for rows.Next() {
		trx := &Trx{}
		err = rows.Scan(&trx.ID, &trx.Name)
		if err != nil {
			return err
		}
		t.data[trx.ID] = trx
	}

	return nil
}


func(ids []int) ([]*ASN, []error) {
  placeholders := make([]string, len(ids))
  args := make([]interface{}, len(ids))
  
  for i := 0; i < len(ids); i++ {
    placeholders[i] = "?"
			args[i] = i
  }
  res := logAndQuery(db,
					"SELECT id, name from dataloader_example.asn WHERE id IN ("+strings.Join(placeholders, ",")+")",
					args...,
				)

  defer res.Close()
  asnById := map[int]*ASN{}

  for res.Next() {
    asn := ASN{}
			err := res.Scan(&asn.ID, &asn.Name)
    if err != nil {
			  panic(err)
			}

    asnById[asn.ID] = &asn
  }
  asns := make([]*ASN, len(ids))
  for i, id := range ids {
    asns[i] = asnById[id]
		  i++
  }

  return asns, nil
}