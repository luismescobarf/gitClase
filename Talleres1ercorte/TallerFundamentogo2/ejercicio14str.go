package main

import "fmt"

func main() {
	var mat [2][5]int
	var mat2 int
	for f := 0; f < 2; f++ {
		for c := 0; c < 5; c++ {
			fmt.Print("Ingrese un numero ")
			fmt.Scan(&mat[f][c])

		}
	}
	fmt.Println("Matriz inicial")
	fmt.Println(mat)
	for c := 0; c < 5; c++ {
		mat2 = mat[0][c]
		mat[0][c] = mat[1][c]
		mat[1][c] = mat2
	}
	fmt.Println("Matriz con cambio")
	fmt.Println(mat)
}
