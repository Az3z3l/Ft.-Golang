package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func mains() {
	graphqlClient := graphql.NewClient("http://127.0.0.1:8080/query")
	graphqlRequest := graphql.NewRequest(`
        {
			hello
        }
    `)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Println(graphqlResponse)
}
