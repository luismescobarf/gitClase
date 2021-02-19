package main
import "fmt"

func main(){
	var num, cont int
	fmt.Println("Ingrese un número de 1 o 2 dígitos")
	fmt.Scan(&num)

	for num!=0{
		num=num/10
		cont++
	}

if (cont==2){
	fmt.Println("El número tiene 2 dígitos")
}else if (cont==1){
	fmt.Println("El número tiene 1 dígito")
}else{
	fmt.Println("El número tiene mas de 2 dígitos")
	}
}