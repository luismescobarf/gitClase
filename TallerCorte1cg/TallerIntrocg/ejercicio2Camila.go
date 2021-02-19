package	main
import "fmt"
func main(){
	var num1, num2, num3, num4 int
		
		fmt.Println("Ingrese 4 numeros ")
		fmt.Scan(&num1)
		fmt.Scan(&num2)
		fmt.Scan(&num3)
		fmt.Scan(&num4)
		
		num2+= num1
		num4*=num3
	
		fmt.Println("La suma de los dos primeros números es:",num2," y La multiplicación de los dos ultimos números es:",num4)
}