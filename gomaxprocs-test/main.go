//use only 1 cpu
package main
import (
	"fmt"
	"runtime"
)


func CpuGoVroooooom(){
	for{

	}
}


func main(){
	runtime.GOMAXPROCS(1)
	fmt.Println("Logical CPUs:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))


	for i:=0;i<100;i++ {
		go CpuGoVroooooom()
	}
	select{}
}

