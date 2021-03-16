package main

import "fmt"

func main() {
	var vec1 [4]int
	var vec2 [4]int
	var vecSuma [4]int
	fmt.Println("Carga del primer vector.")
	for f := 0; f < len(vec1); f++ {
		fmt.Print("Ingrese elemento:")
		fmt.Scan(&vec1[f])
	}
	fmt.Println("Carga del segundo vector.")
	for f := 0; f < len(vec2); f++ {
		fmt.Print("Ingrese elemento:")
		fmt.Scan(&vec2[f])
	}
	for f := 0; f < len(vecSuma); f++ {
		vecSuma[f] = vec1[f] + vec2[f]
	}
	fmt.Println("Vector resultante")
	fmt.Println(vecSuma)
}
