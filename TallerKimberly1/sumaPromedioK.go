package  main 
import "fmt"

func main(){
var n1, n2, n3, n4 float32
fmt.Println("Leer cuatro números enteros. Mostrar la suma y el promedio. ")

fmt.Println("Ingrese el primer número: ")
fmt.Scan(&n1)
fmt.Println("Ingrese el segundo número: ")
fmt.Scan(&n2)
fmt.Println("Ingrese el tercer número: ")
fmt.Scan(&n3)
fmt.Println("Ingrese el cuarto número: ")
fmt.Scan(&n4)
suma := n1+n2+n3+n4
fmt.Println("La suma de los números es: ", suma)

promedio := suma/4
fmt.Println("El promedio de los números es: ", promedio)
}