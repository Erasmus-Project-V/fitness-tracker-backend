package main

import (
	"github.com/pocketbase/pocketbase"
	"log"
)

// zbjx pwal jkcu vzia
func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatalf("Failed to start PocketBase: %v", err)
	}
}
