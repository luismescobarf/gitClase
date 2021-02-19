package main
import "fmt"
func main(){
	var val, cant, abono float32
	fmt.Println("Ingrese valor del articulo")
	fmt.Scan(&val)
	fmt.Println("Ingrese cantidad de articulos")
	fmt.Scan(&cant)
	abono=val*cant
	fmt.Println("El valor a abonar es:", abono)
}