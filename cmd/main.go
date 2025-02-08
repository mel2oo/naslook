package main

import (
	"naslook/internal/app"
	"os"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		panic(err)
	}
}
