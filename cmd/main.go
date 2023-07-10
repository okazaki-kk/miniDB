package main

import (
	"os"

	"github.com/okazaki-kk/miniDB/internal/engine"
	"github.com/okazaki-kk/miniDB/internal/repl"
	"github.com/okazaki-kk/miniDB/storage"
)

func main() {
	catalog := storage.NewCatalog()
	engine := engine.New(*catalog)
	r := repl.New(os.Stdin, os.Stdout, *catalog, *engine)
	r.Start()
}
