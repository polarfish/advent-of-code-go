# advent-of-code-go

All years Advent Of Code solutions in Go

## Prepare

To prepare for solving next puzzle: `go run cmd/prepare.go <year> <day>`.  
Make sure to set environment variable `AOC_SESSION` to download your personal input.

## Solve

To run 2024 day 1: 
```shell
go run cmd/solve.go 2024 1
```  

To run year 2024: 
```shell
go run cmd/solve.go 2024
```  

To run all years:
```shell
go run cmd/solve.go
```

## Solutions

Solutions reside in [solutions](solutions) folder, organized by year.  
Every solution is a separate Go module consisting of a solution file, a test file and an input file.
