package main

import "fmt"

func main() {
	var mat [3][4]int
	var mayor = 0
	for f := 0; f < 3; f++ {
		for c := 0; c < 4; c++ {
			fmt.Print("Ingrese un numero ")
			fmt.Scan(&mat[f][c])
			if mat[f][c] > mayor {
				mayor = mat[f][c]
			}
		}

	}
	fmt.Print("el numero mayor es: ", mayor, "La matriz", mat)
}
