package main

import(	"fmt"
		"time"
	)

func doble(n int) int{ return n * 2 }

func proceso(n int, dobles chan int){
	//Procesar
	n = n * n * n

	//Envío
	dobles<-n
}

func main(){
	
	dobles := make(chan int)
	//n := 5

	go func(){
		for i:=0;i<5;i++{
			time.Sleep(500 * time.Millisecond)
			dobles <- doble(i)
		}
		close(dobles)	
	}()

	for i := range dobles {		
		fmt.Printf("%d ",i)
	}

}