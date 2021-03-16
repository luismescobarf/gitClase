/*Confeccionar una función que recibe un vector de 10 elementos enteros positivos,
retornar la cantidad de elementos con 1 dígito, 2 dígitos y más de dos dígitos.*/
package main
import "fmt"
func digitos(vec [10]int) (int, int, int) {
    uno := 0
    dos := 0
    mas := 0
    for f := 0; f< len(vec); f++ {
        if vec[f] < 10 {
            uno++
        } else {
            if vec[f] <100 {
                dos++
            } else {
                mas++
            }
        }
    }
    return uno, dos, mas
}

func main() {
    vector := [10]int{100, 200, 34, 6, 200, 90, 12, 245, 85, 23}
    v1, v2, v3 := unoDosMasDigitos(vector)
    fmt.Println("Cantidad de elementos con 1 dígito:", v1)
    fmt.Println("Cantidad de elementos con 2 dígitos:", v2)
    fmt.Println("Cantidad de elementos con más de 2 dígitos:", v3)
}