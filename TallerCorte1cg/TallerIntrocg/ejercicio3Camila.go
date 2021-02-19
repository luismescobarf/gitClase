package	main
import "fmt"
func main(){
	var num1, num2, num3, num4, prom float32
		
		fmt.Println("Ingrese 4 numeros ")
		fmt.Scan(&num1)
		fmt.Scan(&num2)
		fmt.Scan(&num3)
		fmt.Scan(&num4)
		
		num4+= num1+num2+num3
		prom=num4/4
	
		fmt.Println("La suma de los dos primeros números es:",num4,prom)
}