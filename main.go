package main

import (
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/vincen320/product-service-graphql-2/cmd"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
