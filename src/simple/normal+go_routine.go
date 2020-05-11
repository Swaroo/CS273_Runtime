package main
import "fmt"
import "time"
import "strconv"

func fun(from string){
  for i:=0; i<=10; i++{
      time.Sleep(100 * time.Millisecond)
      fmt.Println(strconv.Itoa(i)+" from "+from)
  }
}

func main() {
  go fun("Go routine")
  fun("Normal routine")
}
