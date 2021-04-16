package main

import (
	"fmt"
	"time"
)

func doble(n int) int {
	time.Sleep(500 * time.Millisecond)
	return n * 2
}

func triple(n int) int {
	time.Sleep(800 * time.Millisecond)
	return n * 3
}

func main() {

	dobles := make(chan int)
	triples := make(chan int)

	//Generar información
	go func() {

		for i := 0; i < 10; i++ {
			dobles <- doble(i)
			triples <- triple(i)
		}
		close(dobles)
		close(triples)

	}()

	//Consumir la información
	// for d := range dobles {
	// 	fmt.Println(d)
	// }
	// for t := range triples {
	// 	fmt.Println(t)
	// }

lecturaMain:
	for {

		select {
		case i, ok := <-dobles:
			if !ok {
				break lecturaMain
			}
			fmt.Print(i, " ")
		case i, ok := <-triples:
			if !ok {
				break lecturaMain
			}
			fmt.Print(i, " ")

		default:
			/////Código

		}

	}

}
