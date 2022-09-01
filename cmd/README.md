## cmd/karmem

This folder contains files related to the karmem cli tool. The `cmd/karmem` is equivalent of
`protoc` or `flatc`. That is designed to be used as a command line tool, in order to read
karmem files and generate code for the corresponding languages.

## Usage

You can run this program with `go run karmem.org/cmd/karmem help`.

## Structure

- kmcheck:
  - This package will verify if the parsed karmem file (from the `kmparser` package) contains 
any potential conflict, or use any or uses any deprecated features.
- kmgen:
  - This generates code for multiple languages. That will read the parsed 
file (from the `kmparser` package) and generates the code for the given language.
- kmparser:
  - This package is responsible to read karmem schema file. That package is
language independent.
- main.go: The main file for the karmem cli tool (that uses `kmparser`, `kmgen` and `kmcheck`).