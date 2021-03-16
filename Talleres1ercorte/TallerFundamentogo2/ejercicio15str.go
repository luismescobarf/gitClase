package main
import "fmt"
func main() {

  var slice1 []int
	var vl,my int
	suma := 0
	for {
		fmt.Print("Ingrese un entero digite -1 para finalizar:")
		fmt.Scan(&vl)
		if vl == -1 {
			break
		}
		slice1 = append(slice1, vl)
		suma = suma + vl
	}
	prom := suma / len(slice1)

	for f := 0; f < len(slice1); f++ {
		if slice1[f] > prom {
			my++
		}
	}
	fmt.Println("El slice completo es",slice1)
	fmt.Println("Promedio:", prom)
	fmt.Println("La cantidad de valores mayores al promedio son:", my)
}