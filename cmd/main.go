package main

import (
	"os"

	"github.com/014t0/bookshelf/internal"
)

func main() {

	r := internal.Handler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
