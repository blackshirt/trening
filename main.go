package main

import (
	"database/sql"
	log "log"
	http "net/http"
	os "os"

	handler "github.com/99designs/gqlgen/handler"
	"github.com/blackshirt/trening/core/asn"
	"github.com/blackshirt/trening/core/opd"
	"github.com/blackshirt/trening/core/org"
	"github.com/blackshirt/trening/core/trx"
	graph "github.com/blackshirt/trening/graph"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "8080"

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
	opdRepo := opd.NewOpdRepo(conn)
	orgRepo := org.NewOrgRepo(conn)
	asnRepo := asn.NewAsnRepo(conn)
	catRepo := trx.NewTrxCat(conn)
	typeRepo := trx.NewTrxType(conn)
	gqlService := graph.NewRepoServices(asnRepo, opdRepo, orgRepo, catRepo, typeRepo)

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: gqlService})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
