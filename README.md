# coverage-golang-example

This is a basic example to show how we can run the coverage in a golang program?

the `storage_test.go` contains the test for `storage.go` and cover the 100% methods declared.

.
├── README.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── internal
    └── user
        ├── storage.go
        ├── storage_test.go
        └── user.go

# How I run the program?

First, install dependencies
* go mod download

Second, run the main program
* go run cmd/main.go

Finally, run the coverage and inspect it in your browser
* go test -coverprofile=coverage.out ./...
* go tool cover -html=coverage.out
