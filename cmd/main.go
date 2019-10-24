package main

import (
	"log"
	"webexample/internal/app"
)

func main() {
	PORT := 3000

	err := app.Serve(PORT)
	if err != nil {
		log.Fatalf("failed to start app server: %s", err.Error())
	}
}
