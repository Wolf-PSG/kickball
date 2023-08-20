package main

import (
	"go-graphql/main/types"
	"net/http"
  "html/template"
	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
)

type Film struct {
	Title    string
	Director string
}

type landingHandler struct{}

func (h landingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	templ.Execute(w, films)
}

func main() {
	var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: types.RootQuery,
	})

	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

  http.Handle("/", landingHandler{})
	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)

}
	
