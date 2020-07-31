package cmd

import (
	"github.com/graphql-go/graphql"
	handlerGql "github.com/graphql-go/handler"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
)

var serveGql = &cobra.Command{
	Use: "serveGql",
	Short: "Serve your GraphQL server",
	Long: `I've no idea :p'`,
	Run: func(cmd *cobra.Command, args []string) {
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

		h := handlerGql.New(&handlerGql.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: false,
			Playground: true,
		})

		http.Handle("/graphql", h)
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func init()  {
	rootCmd.AddCommand(serveGql)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func hostname() string{
	hostname, _ := os.Hostname()
	return hostname
}