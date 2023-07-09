package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"

	"github.com/okazaki-kk/miniDB/internal/parser"
	"github.com/okazaki-kk/miniDB/internal/parser/lexer"
)

const PROMPT = "miniDB >>"

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello " + user.Name + "!")
	fmt.Println("This is the miniDB!")
	fmt.Println("Feel free to type in commands")
	Start()
}

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		p := parser.New(lexer.New(line))
		stmts, err := p.Parse()

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(stmts)
	}
}
