package main

import("fmt")

func doble(n int) int{ return n * 2 }

func main(){
	
	var ch chan int	
	var a int
	var contenedor []int
	fmt.Printf(" %T %v \n ", ch, ch)
	fmt.Printf(" Como una variable-> %T %v \n ", a, a)
	fmt.Printf(" Como un slice-> %T %v \n ", contenedor, contenedor)
	
}