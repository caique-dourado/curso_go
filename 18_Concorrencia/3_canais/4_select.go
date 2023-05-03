package main

import (
	"fmt"
	"time"
)

func main() {
	ch, ch2 := make(chan int), make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for j := 1; j < 10; j++ {
			ch2 <- j
			time.Sleep(time.Second * 1)
		}
		close(ch2)
	}()

	for k := 1; k < 20; k++ {
		select {
		case chReturn := <-ch:
			fmt.Println("Retorno do ch:", chReturn)
		case ch2Return := <-ch2:
			fmt.Println("Retorno do ch2:", ch2Return)
		}
	}

}
