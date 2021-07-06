package main

import (
	"fmt"
	"lesson8/foo"
	"os"
)

func main()  {
	conf := foo.NewConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	diagPort := os.Getenv("DIAG_PORT")
	if diagPort == "" {
		diagPort = "8081"
	}

	name := os.Getenv("NAME")
	if name == "" {
		name = "GO student"
	}

	id := os.Getenv("ID")
	if id == "" {
		id = "0"
	}

	fmt.Println(conf)
}