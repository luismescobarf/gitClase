package main

import "fmt"

func main() {
	var mat [3][4]int
	for f := 0; f < 3; f++ {

		for c := 0; c < 4; c++ {
			fmt.Print("Ingrese un numero ", f, c, " ")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Print(mat)
	fmt.Print("Primela fila ")
	for c := 0; c < 4; c++ {

		fmt.Print(mat[0][c], "-")
	}
	fmt.Print("Ultima fila ")
	for c := 0; c < 4; c++ {

		fmt.Print(mat[2][c], "-")
	}
	fmt.Print("Primera Columna ")
	for f := 0; f < 3; f++ {

		fmt.Print(mat[f][0], "-")
	}

}
