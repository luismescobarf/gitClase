package  main 
import "fmt"

func main(){
var n1, n2, n3 float32
fmt.Println("Se ingresan tres notas si el promedio es mayor o igual a siete mostrar mensaje 'Promocionado' ")

fmt.Println("Ingrese el primer número: ")
fmt.Scan(&n1)
fmt.Println("Ingrese el segundo número: ")
fmt.Scan(&n2)
fmt.Println("Ingrese el segundo número: ")
fmt.Scan(&n3)

suma := n1+n2+n3
promedio := suma/3

if promedio>=7 {
	fmt.Println("Promocionado")
}else{
	fmt.Println("Repitente")
}
}