// 2. Obtener el mayor elemento de una matriz bidimensional de números enteros.

package main

import (
	"fmt"
	"sync"

	"time"
)

var inicio time.Time

var fin time.Time

var matriz [3][3]int

var aux = 0

var wg sync.WaitGroup

func llenarMatriz() {

	for i := 0; i < 3; i++ {

		for j := 0; j < 3; j++ {

			fmt.Println("Ingrese el número: ")

			fmt.Scan(&matriz[i][j])

		}

	}

}

func mayorFila(fila [3]int, mayores *[3]int, numFila int) {

	defer wg.Done()

	var mayorFila int

	for j := 0; j < 3; j++ {

		if fila[j] > mayorFila {

			mayorFila = fila[j]

		}

	}

	mayores[numFila] = mayorFila

	fmt.Println("Mayor fila ", mayorFila)

	//return mayorFila

}

func mostrarMayor() {

	var mayores [3]int

	for i := 0; i < 3; i++ {

		wg.Add(1)
		go mayorFila(matriz[i], &mayores, i)

	}
	wg.Wait()

	for i := 0; i < 3; i++ {

		if mayores[i] > aux {
			aux = mayores[i]
		}

	}

}

func main() {

	//1 2 3
	//4 5 6
	//7 8 9

	inicio = time.Now()

	// fmt.Println(inicio.Second())

	llenarMatriz()

	mostrarMayor()

	fmt.Println("El número mayor es: ")

	fmt.Println(aux)

	fin = time.Now()

	// fmt.Println(fin.Second())

	diff := fin.Sub(inicio)

	fmt.Println("Tiempo de ejecución: ")

	fmt.Println(diff)

}
