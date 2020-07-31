package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	type Employee struct {
		Name string `json:"name"`
		Address string `json:"address"`
	}

	var employees = []Employee{
		{
			Name:    "Arif",
			Address: "Jakarta",
		},
		{
			Name:    "Teguh",
			Address: "Jakarta",
		},
	}

	var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"name": &graphql.Field{Type: graphql.String},
			"address": &graphql.Field{Type: graphql.String},
		},
	})

	var QueryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employees": &graphql.Field{
				Type: graphql.NewList(EmployeeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return employees, nil
				},
			},
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "hello world", nil
				},
			},
		},
	})


	//rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: QueryType}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
