package main

import (
	"os"

	"go-account-service/application"
)

func main() {
	port := os.Args[1]
	application.Run(port)
}
