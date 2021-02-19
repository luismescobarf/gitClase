package main
import "fmt"
func main(){
	var num1, num2, op1, op2 float32
	fmt.Println("Ingrese 2 números")
	fmt.Scan(&num1)
	fmt.Scan(&num2)

	if (num1>num2){
		op1=num1+num2
		op2=num1-num2
		fmt.Println("El número",num1,"es mayor a el número", num2, "la suma es:",op1,"La resta es:",op2)
	} else {
		op1=num1*num2
		op2=num1/num2
		fmt.Println("El número",num2,"es mayor  a el número", num1, "la multiplicacion es:",op1,"La division es:",op2)
	}
}