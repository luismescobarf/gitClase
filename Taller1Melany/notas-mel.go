package  main 
import "fmt"

func main(){
var n1, n2, n3 float32
fmt.Print("Ingrese la primera nota: ")
fmt.Scan(&n1)
fmt.Print("Ingrese la segunda nota: ")
fmt.Scan(&n2)
fmt.Print("Ingrese la tercera nota: ")
fmt.Scan(&n3)
suma :=n1+n2+n3
promedio := suma/3

if promedio>=7{
	
fmt.Println("Promocionado")

}else {
	fmt.Println("Repitente")
}


}