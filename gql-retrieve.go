package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func main() {
	graphqlClient := graphql.NewClient("http://127.0.0.1:5555/query")

	req := graphql.NewRequest(`
	query ppp($a: String! $b: String!){
		a:	ping(name: $a)  
		b:	ping(name: $b)	
	}  
	`)

	// var variables = "adsa"
	req.Var("a", "asd")
	req.Var("b", "asd")
	fmt.Println(req)
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), req, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Println(graphqlResponse)
}
