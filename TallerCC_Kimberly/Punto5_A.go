// 5. Implementar una solución para calcular el producto de dos matrices. Diseñar e
// implementar una solución distribuída de este procedimiento.
package main

import ("fmt"
"time"
)

var inicio time.Time
var fin time.Time

var matrizA [3]int
var matrizB [3]int
var matrizC [3]int

func  llenarMatriz()  {
    for i :=0 ; i < 3; i++ {
        
        fmt.Println("Ingrese el número para la matriz A: ")
        fmt.Scan(&matrizA[i])

		fmt.Println("Ingrese el número para la matriz B: ")
        fmt.Scan(&matrizB[i])
    }
}

func  producto()  {
	for i:=0; i<3;i++{
		matrizC[i]=matrizA[i]*matrizB[i]
			
	}	
}

func main()  {
	llenarMatriz()
	inicio = time.Now()
	// fmt.Println(inicio.Second())

	producto()
	
	fmt.Println("Producto cartesiano ")
	fmt.Println(matrizC)
	
	fin = time.Now()
	// fmt.Println(fin.Second())

	diff := fin.Sub(inicio)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(diff)
}