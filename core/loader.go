package loader

import (
	"log"
	"strings"

	"github.com/blackshirt/trening/models"
)

func Gets(ids []int) ([]*models.ASN, []error) {
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i := 0; i < len(ids); i++ {
		placeholders[i] = "?"
		args[i] = i
	}
	/*
		res := logAndQuery(db,
			"SELECT id, name from user WHERE id IN ("+
				strings.Join(placeholders, ",")+")",
			args...,
		)
	*/
	placeholders = strings.Join(placeholders, ",")
	query := `SELECT id, name, nip, current_job, current_grade, current_places FROM asn WHERE id IN "+placeholders+"`
	defer res.Close()

	asns := make([]*models.ASN, len(ids))
	i := 0
	for res.Next() {
		asns[i] = &models.ASN{}
		err := res.Scan(
			&asns[i].ID,
			&asns[i].Name,
			&asns[i].Nip,
			&asns[i].CurrentJob,
			&asns[i].CurrentGrade,
			&asns[i].CurrentPlaces,
		)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

	return asns, nil
}
