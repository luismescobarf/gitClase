package main

import("fmt")

func doble(n int) int{ return n * 2 }

func main(){
	
	ch := make(chan chan int)
	//var ch chan int
	
	fmt.Printf(" %T \n ", ch)
	
}