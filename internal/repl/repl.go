package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"

	"github.com/okazaki-kk/miniDB/internal/sql"
)

const PROMPT = "miniDB >> "

type Repl struct {
	input    io.Reader
	output   io.Writer
	catalog  sql.Catalog
	database sql.Database
}

func New(input io.Reader, output io.Writer, catalog sql.Catalog) *Repl {
	return &Repl{
		input:   input,
		output:  output,
		catalog: catalog,
	}
}

func (r Repl) Start() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	io.WriteString(r.output, fmt.Sprintf("Hello %s!\n", user.Name))
	io.WriteString(r.output, "This is the miniDB!\n")
	io.WriteString(r.output, "Feel free to type in commands\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		io.WriteString(r.output, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if message, err := r.execCommand(line); err != nil {
			io.WriteString(r.output, fmt.Sprintf("%s\n", err.Error()))
		} else {
			io.WriteString(r.output, message)
		}
	}
}

func (r *Repl) execCommand(input string) (string, error) {
	cmd := strings.TrimSpace(input)
	params := strings.Fields(cmd)

	switch params[0] {
	case `\use`:
		return r.useDatabase(params)
	default:
		return "", fmt.Errorf("unknown command: %v", params[0])
	}
}

func (r *Repl) useDatabase(params []string) (string, error) {
	if len(params) < 2 {
		return "", fmt.Errorf("database name not specified")
	}

	db, err := r.catalog.GetDatabase(params[1])
	if err != nil {
		return "", err
	}

	r.database = db
	io.WriteString(r.output, fmt.Sprintf("database %s using", db.Name()))

	return "database changed\n", nil
}
