package main

import (
	"fmt"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nickmartinwebdev/graphql-go/graphql/resolver"
	"github.com/nickmartinwebdev/graphql-go/graphql/schema"
	"github.com/rs/cors"
)

func main() {

	fmt.Println("starting program")

	s, err := schema.String()
	if err != nil {
		panic("error creating schema")
	}

	root, _ := resolver.NewRoot()
	schema := graphql.MustParseSchema(s, root)

	mux := http.NewServeMux()
	mux.Handle("/graphql", &relay.Handler{Schema: schema})
	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":4000", handler)
}
