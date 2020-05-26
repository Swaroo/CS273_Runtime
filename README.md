# Analysis of Golang Concurrency Model
Course Project for Runtime Systems - Spring 2020

#### For the initial Go-routine programs in src/simple
To run : go run <program.go>

### How nomal functions and go-routine run simultaneously
go run src/simple/normal+go_routine.go

### Multiplication through Go-routine and channels
go run src/simple/channels_1.go

### Equivalent Binary Trees through Go-routines and channels
Packages - go get golang.org/x/tour/tree
To run - go run src/advanced/equivalent_binary_trees.go

#### To benchmark Standalone Go programs
Inside the directory containing the Go programs:
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
