
package main

import ("fmt"
"time"
)

var inicio time.Time
var fin time.Time

var matriz [4][4]int
var aux = 0

func  llenarMatriz()  {
    for i :=0 ; i < 4; i++ {
        for j := 0; j < 4; j++ {
            fmt.Println("Ingrese el número: ")
            fmt.Scan(&matriz[i][j])
        }
    }
}

func  mostrarMayor()  {
	for i:=0; i<4;i++{
		for j:=0;j<4;j++{
			if(matriz[i][j]>aux){
				aux=matriz[i][j]
			}
		}
	}	
}

func main()  {
	inicio = time.Now()
	// fmt.Println(inicio.Second())

	go llenarMatriz()
	go mostrarMayor()
	
	fmt.Println("El número mayor es: ")
	fmt.Println(aux)
	
	fin = time.Now()
	// fmt.Println(fin.Second())

	diff := fin.Sub(inicio)
	fmt.Println("Tiempo de ejecución: ")
	fmt.Println(diff)
}