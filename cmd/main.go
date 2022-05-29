package main

import (
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/nickmartinwebdev/graphql-go/graphql/resolver"
	"github.com/nickmartinwebdev/graphql-go/graphql/schema"
)




func main () {

	fmt.Println("starting program")

	s, err := schema.String()
	if err != nil {
		panic("error creating schema")
	}

	root, _ := resolver.NewRoot()
    schema := graphql.MustParseSchema(s, root)
    http.Handle("/graphql", &relay.Handler{Schema: schema})
    log.Fatal(http.ListenAndServe(":8080", nil))
}
