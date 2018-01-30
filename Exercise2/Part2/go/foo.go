// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
)

// Control signals
//const (
//	GetNumber = iota
//	Exit
//)

func number_server(add_number <-chan int, number chan<- int, finished <- chan bool) {
	var i = 0
	var done = 0

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		select {
            //case mes1 := <=
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
            case j := <-add_number:
                i=i+j
            case <- finished:
                if done == 0 {
                    done =1
                    
                }   else {
                    number<-i
                }
        }
    }
                
}

func incrementing(add_number chan<-int, finished chan<- bool) {
	for j := 0; j<1000000; j++ {
		add_number <- 1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func decrementing(add_number chan<- int, finished chan<- bool) {
	for k := 0; k<1000000; k++ {
		add_number <- -1
	}
	//TODO: signal that the goroutine is finished
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO: Construct the required channels
	// Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
    add_number := make(chan int,2)
    finished := make(chan bool)
    number :=  make(chan int)

	// TODO: Spawn the required goroutines
    go number_server (add_number,number,finished)
    go incrementing(add_number,finished)
    go decrementing(add_number,finished)

	// TODO: block on finished from both "worker" goroutines

	//control<-GetNumber
	Println("The magic number is:", <- number)
	//control<-Exit
}
