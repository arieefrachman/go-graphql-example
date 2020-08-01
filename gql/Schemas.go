package gql

import "github.com/graphql-go/graphql"

var PersonType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Person",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.Int},
		"first_name": &graphql.Field{Type: graphql.String},
		"last_name": &graphql.Field{Type: graphql.String},
		"birth_date": &graphql.Field{Type: graphql.String},
	},
})
