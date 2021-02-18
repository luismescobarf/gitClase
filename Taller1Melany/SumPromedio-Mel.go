package  main 
import "fmt"

func main () {
var n1, n2, n3, n4 float32
fmt.Print("Ingrese el primer numero entero: ")
fmt.Scan(&n1)
fmt.Print("Ingrese el segundo numero entero: ")
fmt.Scan(&n2)
fmt.Print("Ingrese el tercer numero entero: ")
fmt.Scan(&n3)
fmt.Print("Ingrese el cuarto numero entero: ")
fmt.Scan(&n4)
suma := n1+n2+n3+n4
fmt.Println("La suma de los numeros es: ", suma)
promedio := suma/4
fmt.Println("El promedio de los numeros es: ", promedio)
}