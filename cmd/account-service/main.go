package main

import (
	"os"

	"go-account-service/application"
)

func main() {
	application.Run(os.Args[1])
}
