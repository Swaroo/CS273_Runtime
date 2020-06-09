# Analysis of Golang Concurrency Model
Course Project for Runtime Systems - Spring 2020

#### For CPU Bound task
cd benchmarking/CPU_bound-multiply
go test -cpu=4 -bench=.

#### For IO Bound task
cd benchmarking/IO_bound-search
##### To generate text files (Creates 50 files)
cd benchmarking/IO_bound-search/search_directory
python3 random_generator.py
##### To evaluate
go test -cpu=4 -bench=.

#### For CPU + I/O Bound task
cd benchmarking/URL_load_json_unmarshall
go test -cpu=4 -bench=.

#### CPU parameter
All the above benchmarks can be run with -cpu=1 / 2 / 3 /4 till the number of physical CPU Cores.


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
