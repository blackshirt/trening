package loader

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/blackshirt/trening/models"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type loaders struct {
	OpdId *OPDLoader
}

func LoaderMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		// set this to zero what happens without dataloading
		wait := 250 * time.Microsecond

		// simple 1:1 loader, fetch an address by its primary key
		ldrs.OpdId = &OPDLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []int) ([]*models.OPD, []error) {
				placeholders := make([]string, len(keys))
				args := make([]interface{}, len(keys))
				for i := 0; i < len(keys); i++ {
					placeholders[i] = "?"
					args[i] = i
				}
				stmt, err := db.Prepare("SELECT id, name, long_name, road_number, city, province from opd WHERE id IN (" + strings.Join(placeholders, ",") + ")")

				if err != nil {
					log.Fatal(err)
				}
				rows, err := stmt.Query(args...)
				if err != nil {
					log.Fatal(err)
				}
				defer rows.Close()

				opdById := map[int]*models.OPD{}
				for rows.Next() {
					opd := models.OPD{}
					err := rows.Scan(&opd.ID, &opd.Name, &opd.LongName,
						&opd.RoadNumber, &opd.City, &opd.Province)
					if err != nil {
						panic(err)
					}
					opdById[opd.ID] = &opd
				}

				opds := make([]*models.OPD, len(keys))
				for i, id := range keys {
					opds[i] = opdById[id]
					i++
				}

				return opds, nil
			},
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
