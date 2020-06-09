# Analysis of Golang Concurrency Model
Course Project for Runtime Systems - Spring 2020

### Installing Golang
https://golang.org/dl/ <br>
Follow the instructions from the official site to obtain Golang for your specific OS. <br>
No extra pacakages are required to run these programs.

#### Benchmarking CPU Bound task
cd benchmarking/CPU_bound-multiply <br>
go test -cpu=4 -bench=.

#### Benchmarking IO Bound task
cd benchmarking/IO_bound-search
##### Generate text files (Creates 50 files) for temporary search
cd benchmarking/IO_bound-search/search_directory<br>
python3 random_generator.py
##### Evaluation
go test -cpu=4 -bench=.

#### Benchmarking CPU + I/O Bound task
cd benchmarking/URL_load_json_unmarshall<br>
go test -cpu=4 -bench=.

#### CPU parameter
All the above benchmarks can be run with -cpu=1 / 2 / 3 /4 till the number of physical CPU Cores.<br>
The value specifies the number of virtual cores offered to the benchmarking program.


#### For the initial Go-routine programs in src/simple
To run : go run <program.go>

### How nomal functions and go-routine run simultaneously
go run src/simple/normal+go_routine.go

### Multiplication through Go-routine and channels
go run src/simple/channels_1.go

### Equivalent Binary Trees through Go-routines and channels
Packages - go get golang.org/x/tour/tree <br>
To run - go run src/advanced/equivalent_binary_trees.go

#### To benchmark Standalone Go programs
Inside the directory containing the Go programs: <br>
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
