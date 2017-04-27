package main

import (
	"github.com/AKovalevich/dapper/routes"

	"os"
	"fmt"
)

func main() {
	os.Exit(run())
}

func run() (int) {
	route := &routes.Route{
		Structure: func () (interface{}, error) {
			return "a", nil
		},
	}

	fmt.Print(route.Structure())
	return 1
}
