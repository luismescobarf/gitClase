
package main

import ("fmt"
"time"
)

var inicio time.Time
var fin time.Time

var matriz1 [3]int
var matriz2 [3]int
var matriz3 [3]int
var aux= 0

func  llenarMatriz()  {
    for i :=0 ; i < 3; i++ {
        
            fmt.Println("Ingrese el número para la primera matriz: ")
            fmt.Scan(&matriz1[i])
			fmt.Println("Ingrese el número para la segunda matriz: ")
            fmt.Scan(&matriz2[i])
			
        
    }
}

func  producto()  {
	for i:=0; i<3;i++{
	
			matriz3[aux]=matriz1[i]*matriz2[i]
			aux++
		
	}	
}

func main()  {
	llenarMatriz()
	inicio = time.Now()
	// fmt.Println(inicio.Second())

	producto()
	
	fmt.Println("El resultado es: ")
	fmt.Println(matriz3)
	
	fin = time.Now()
	// fmt.Println(fin.Second())

	diff := fin.Sub(inicio)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(diff)
}