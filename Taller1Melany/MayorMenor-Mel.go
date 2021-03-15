package  main 
import "fmt"

func main(){
var n1, n2 float32
fmt.Print("Ingrese el primer numero: ")
fmt.Scan(&n1)
fmt.Print("Ingrese el segundo numero: ")
fmt.Scan(&n2)
if n1>n2{
	suma := n1+n2
fmt.Println("La suma de los numeros es: ", suma)
resta := n1-n2
fmt.Println("La diferencia de los numeros es: ", resta)

}else {
	multi := n1*n2
fmt.Println("La multiplicacion de los numeros es: ", multi) 
	division := n1/n2
	fmt.Println("La division de los numeros es: ", division)
}


}