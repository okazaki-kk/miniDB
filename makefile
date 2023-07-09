fmt:
	go fmt ./...

test:
	go test ./... -v

test-pretty:
	set -o pipefail && go test -v ./... fmt -json | tparse -all

repl:
	go run ./cmd/main.go
