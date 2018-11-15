package main

import (
	"context"
	"database/sql"
	log "log"
	http "net/http"
	os "os"
	"time"

	handler "github.com/99designs/gqlgen/handler"
	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/core/opd"
	"github.com/blackshirt/trening/core/org"
	graph "github.com/blackshirt/trening/graph"
	"github.com/blackshirt/trening/models"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"asnCtx"}

type loaders struct {
	asnByID *models.ASNLoader
}

func LoaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		wait := 250 * time.Microsecond
		ldrs.asnByID = &models.ASNLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []int) ([]*models.ASN, []error) {

			},
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func ctxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn, err := sql.Open("mysql", "root:123@/trainix")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	opdRepo := opd.NewOPDRepo(conn)
	orgRepo := org.NewOrgRepo(conn)
	asnRepo := asn.NewASNRepo(conn)
	gqlService := graph.NewGraphQLService(asnRepo, opdRepo, orgRepo)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", LoaderMiddleware(handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: gqlService}))))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
