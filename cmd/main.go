package main

import (
	"os"

	"github.com/okazaki-kk/miniDB/internal/repl"
	"github.com/okazaki-kk/miniDB/storage"
)

func main() {
	r := repl.New(os.Stdin, os.Stdout, storage.NewCatalog())
	r.Start()
}
