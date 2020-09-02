package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/machinebox/graphql"
)

func graphiql(w http.ResponseWriter, r *http.Request) {
	graphqlClient := graphql.NewClient("http://127.0.0.1:5001/query")
	graphqlRequest := graphql.NewRequest(`
		{
			person{
				id
				name
				pet{
					id
					name
				}
			}
		}
    `)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, graphqlResponse)
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func main() {
	http.HandleFunc("/", graphiql)
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":5000", nil)
}
