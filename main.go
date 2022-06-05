package main

import (
	"go-graphql/main/types"
	"net/http"

	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: types.RootQuery,
	})

	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)

}
