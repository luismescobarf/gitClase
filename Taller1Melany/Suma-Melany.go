package  main 
import "fmt"

func main(){
var n1, n2, n3, n4 int
fmt.Print("Ingrese el primer numero: ")
fmt.Scan(&n1)
fmt.Print("Ingrese el segundo numero: ")
fmt.Scan(&n2)
fmt.Print("Ingrese el tercer numero: ")
fmt.Scan(&n3)
fmt.Print("Ingrese el cuarto numero: ")
fmt.Scan(&n4)
suma := n1+n2
fmt.Println("La suma de los primeros nuemeros es: ", suma)
multi := n3*n4
fmt.Println("La multiplicacion de los ultimos numeros es: ", multi)
}

