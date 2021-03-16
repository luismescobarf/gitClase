package main

import "fmt"

func main() {
	var mat [2][5]int
	for f := 0; f < 2; f++ {
		for c := 0; c < 5; c++ {
			fmt.Print("Ingrese elemento:")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Println("Matriz completa")
	fmt.Println(mat)
	var aux int
	for c := 0; c < 5; c++ {
		aux = mat[0][c]
		mat[0][c] = mat[1][c]
		mat[1][c] = aux
	}
	fmt.Println("La matriz completa luego de intercambiar primer fila con la segunda.")
	fmt.Println(mat)
}
