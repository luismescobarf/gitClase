package main

import "fmt"

func digitos(vector [10]int) (int, int, int) {
  
	d1 := 0
	d2 := 0
	d3 := 0

	for i := 0; i < len(vector); i++ {
		if vector[i]<10 && vector[i]>-10 {
			d1++		
		}else{
			if vector[i]<100&&vector[i]>-100{
				d2++
			}else{
				d3++
			}
		}
	}
	fmt.Println(vector)
	
  return d1, d2, d3
}


func main() {
    var array[10] int
	
	for i := 0; i < 10; i++ {
		fmt.Println("Ingrese un número entero: ")
		fmt.Scan(&array[i])
	}

	v1, v2, v3 := digitos(array)
    
    fmt.Println("\nCantidad de elementos con 1 dígito:", v1)
    fmt.Println("Cantidad de elementos con 2 dígitos:", v2)
    fmt.Println("Cantidad de elementos con más de 2 dígitos:", v3)
}