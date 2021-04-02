package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"imkernel.dev/logos"
	"imkernel.dev/logos/graph"
	"imkernel.dev/logos/graph/generated"
	"imkernel.dev/logos/internal/app"
)

const defaultPort = "8080"

func main() {

	f, err := logos.EmbeddedBible.Open("data/en/kjv.xml")
	// f, err := os.Open("./data/en/kjv.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	startedParsingOSISFile := time.Now()
	bible, err := app.ParseOSISFile(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Parsed the %s Bible from the osis file successfully!\n", bible.Identifier)
	fmt.Printf("Time taken to parse: %s\n", time.Since(startedParsingOSISFile))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
