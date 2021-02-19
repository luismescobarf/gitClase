package main
import "fmt"
func main(){
	var not1,not2,not3,prom float32
	fmt.Println("Ingrese notas")
	fmt.Scan(&not1)
	fmt.Scan(&not2)
	fmt.Scan(&not3)

	 prom=(not1+not2+not3)/3

	if (prom>=7){
		fmt.Println("Promocionado")
	}

}