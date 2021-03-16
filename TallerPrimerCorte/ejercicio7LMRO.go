package main

import "fmt"

func main() {
	var alturas [5]float32
	var suma float32
	for f := 0; f < 5; f++ {
		fmt.Print("Ingrese la altura de la persona:")
		fmt.Scan(&alturas[f])
		suma = suma + alturas[f]
	}
	promedio := suma / 5
	fmt.Println("Promedio de las alturas:", promedio)
	may := 0
	men := 0
	for f := 0; f < 5; f++ {
		if alturas[f] > promedio {
			may++
		} else {
			if alturas[f] < promedio {
				men++
			}
		}
	}
	fmt.Println("Cantidad de personas mayores al promedio:", may)
	fmt.Println("Cantidad de personas menores al promedio:", men)
}
