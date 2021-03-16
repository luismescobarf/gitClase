package main

import "fmt"

func main() {
	var mat [4][4]int
	for f := 0; f < 4; f++ {
		for c := 0; c < 4; c++ {
			fmt.Print("Ingrese componente:")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Println("Matriz completa.")
	fmt.Println(mat)
	fmt.Println("Elementos de la diagonal principal.")
	for k := 0; k < 4; k++ {
		fmt.Print(mat[k][k], " ")
	}
}
