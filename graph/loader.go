package graph

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/blackshirt/trening/models"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"userCtx"}

type loaders struct {
	OpdId *ASNLoader
}

func LoaderMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}

func asnFetchFn() func([]string) ([]*models.Asn, []error) {
	return func(keys []string) (results []*models.Asn, errors []error) {
		results = make([]*model.Asn, len(keys))
		errors = make([]error, len(keys))

		return
	}
}
