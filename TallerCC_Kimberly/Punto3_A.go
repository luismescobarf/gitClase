// 3. Generar el conjunto resultante del producto cartesiano de dos conjuntos
// ingresados. Los conjuntos estarán representados en vectores (slices o arrays).
package main

import ("fmt"
"time"
)

var inicio time.Time
var fin time.Time

var matrizA [3]int
var matrizB [3]int
var matrizC [9]int
var aux = 0

func  llenarMatriz()  {
    for i :=0 ; i < 3; i++ {
        
        fmt.Println("Ingrese el número para la matriz A: ")
        fmt.Scan(&matrizA[i])

		fmt.Println("Ingrese el número para la matriz B: ")
        fmt.Scan(&matrizB[i])
    }
}

func  productoCartesiano()  {
	for i:=0; i<3;i++{
		for j:=0;j<3;j++{
			matrizC[aux]=matrizA[i]*matrizB[j]
			aux++
		}
	}	
}

func main()  {
	llenarMatriz()
	inicio = time.Now()
	// fmt.Println(inicio.Second())

	productoCartesiano()
	
	fmt.Println("Producto cartesiano ")
	fmt.Println(matrizC)
	
	fin = time.Now()
	// fmt.Println(fin.Second())

	diff := fin.Sub(inicio)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(diff)
}