package main

import "fmt"

func main() {
	var mat [4][4]int
	for f := 0; f < len(mat); f++ {

		for c := 0; c < len(mat); c++ {
			fmt.Print("Ingrese un numero", f, c, " ")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Print(mat)

	for k := 0; k < 4; k++ {
		fmt.Print(mat[k][k], "Diagonal", "")
	}
}
