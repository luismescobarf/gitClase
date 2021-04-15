package main

import ("fmt"
"time"
)

var inicio time.Time
var fin time.Time
func  mostrar0()  {
	
	for i := 1; i <= 1000; i++{
		fmt.Print("0 - ")
		
	}	
	
}

func  mostrar1()  {
	for i := 1; i <= 1000; i++{
		fmt.Print("1 - ")
		
	}	
	
}

func main()  {
	inicio = time.Now()
	fmt.Println(inicio.Second())

	go mostrar0()
	go mostrar1()
	
	
	fin = time.Now()
	fmt.Println(fin.Second())

	diff := fin.Sub(inicio)
	
	var continua string
    fmt.Scan(&continua)

	fmt.Println(diff)
}