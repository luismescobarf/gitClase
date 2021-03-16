package main
import "fmt"
import "sort"

func main(){
	var l,n string
	fmt.Println("Ingrese dos nombres de personas")
	fmt.Scan(&l)
	fmt.Scan(&n)
	strs := []string{l, n}
    sort.Strings(strs)
    fmt.Println("Orden alfabetico", strs)
}