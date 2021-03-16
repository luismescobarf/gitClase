package main

import "fmt"

func main() {
	var mat [3][4]int
	for f := 0; f < 3; f++ {
		for c := 0; c < 4; c++ {
			fmt.Print("Ingrese elemento:")
			fmt.Scan(&mat[f][c])
		}
	}
	fmt.Println("Matriz completa")
	fmt.Println(mat)
	may := mat[0][0]
	for f := 0; f < 3; f++ {
		for c := 0; c < 4; c++ {
			if mat[f][c] > may {
				may = mat[f][c]
			}
		}
	}
	fmt.Print("El elemento mayor de la matriz es: ", may)
}
