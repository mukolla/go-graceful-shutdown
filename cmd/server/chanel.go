package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start main.")

	//data := make(chan int)
	//exit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-data)
	//	}
	//
	//	exit <- 0
	//}()
	//selectOne(data, exit)

	ch := make(chan int)
	go sayHello(ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println("Completed main.")
}

func say(word string) {
	time.Sleep(1 * time.Second)
	fmt.Println(word)
}

func sayHello(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		say("hello")
		exit <- i
	}

	exit <- 17

	close(exit)
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
			fmt.Println("data ...")
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println(".... waiting...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
