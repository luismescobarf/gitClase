package main

import("fmt")

func doble(n int) int{ return n * 2 }

func proceso(n int, dobles chan int){
	//Procesar
	n = n * n * n

	//Envío
	dobles<-n
}

func main(){
	
	dobles := make(chan int)
	n := 5

	go func(){
		dobles <- doble(n)
	}()

	go proceso(n, dobles)

	fmt.Println("Primera lectura recibida por el canal:", <- dobles)
	fmt.Println("Segunda lectura recibida por el canal:", <- dobles)

	/*
	valorBooleano := func(a int, b int) bool{
		c := a + b
		fmt.Println(c)
		return true
	}(6,7)
	
	fmt.Print(valorBooleano)
	*/
	
	
	
	
}