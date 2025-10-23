# advent-of-code-go

All years Advent Of Code solutions in Go

## Prepare

To prepare for solving next puzzle: `go run cmd/prepare/prepare.go <year> <day>`.  
Make sure to set environment variable `AOC_SESSION` to download your personal input.

## Solve

Run one day: 
```shell
go run cmd/solve/solve.go 2025 1
```  

Run full year: 
```shell
go run cmd/solve/solve.go 2015
```  

Run everything:
```shell
go run cmd/solve/solve.go
```

## Solutions

Solutions reside in [solutions](solutions) folder, organized by year.  
Every solution is a separate Go module consisting of a solution file, a test file and an input file.
