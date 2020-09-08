package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func main() {
	graphqlClient := graphql.NewClient("http://127.0.0.1:5555/query")
	req := graphql.NewRequest(`
	mutation register($in:register_input){
		register(input:$in)
	}
	`)
	data := map[string]string{
		"email":    "Dog",
		"username": "Cat",
		"password": "Cow",
	}
	req.Var("in", data)
	fmt.Println(req)

	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), req, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Println(graphqlResponse)
}
