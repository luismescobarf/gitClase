package  main 
import "fmt"

func main(){
var n1, n2, n3, n4 int
fmt.Println("Leer cuatro números y mostras la suma de los dos primeros y la multiplicación de los dos últimos. ")

fmt.Println("Ingrese el primer número: ")
fmt.Scan(&n1)
fmt.Println("Ingrese el segundo número: ")
fmt.Scan(&n2)
fmt.Println("Ingrese el tercer número: ")
fmt.Scan(&n3)
fmt.Println("Ingrese el cuarto número: ")
fmt.Scan(&n4)
suma := n1+n2
fmt.Println("La suma de los primeros números es: ", suma)
multi := n3*n4
fmt.Println("La multiplicación de los últimos números es: ", multi)
}