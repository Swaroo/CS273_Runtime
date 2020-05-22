package main
import "runtime/pprof"
import "flag"
import "os"
import "runtime"
import "log"
import "fmt"
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func fibo(n int) int{
	i,j := 0,1
	for k:=2; k<n; k++{
		i,j = j, i+j 
	}
	return j
}

func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

	fmt.Println("Starting pprof")
	out := fibo(5)
	fmt.Println(out)
	out = fibo(50)
	fmt.Println(out)
	out = fibo(100)
	fmt.Println(out)

    if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}
