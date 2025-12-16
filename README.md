# advent-of-code-go

All years Advent Of Code solutions in Go

## Solutions

Check [solutions](solutions) folder (organized by year).  
Every solution is a separate Go module consisting of a solution file, a test file and an input file.

## Prepare

To prepare for solving next puzzle: `go run ./cmd/prepare <year> <day>`.  
Make sure to set environment variable `AOC_SESSION` to download your personal input.

## Inputs

Input files are not included in the repository (ignored from the version control) as requested in [FAQ](https://adventofcode.com/2025/about#faq_copying).  
Your personal inputs are downloaded automatically when you use `prepare` command.  
You can also download all missing inputs with `input` command:   
```shell
go run ./cmd/input
```

## Solve

Run one day: 
```shell
go run ./cmd/solve 2025 1
```  

Run full year: 
```shell
go run ./cmd/solve 2015
```  

Run everything:
```shell
go run ./cmd/solve
```

