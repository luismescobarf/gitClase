package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	c, err := net.Dial("tcp", "localhost:12345")
	if err != nil {

		fmt.Println(err)
		return

		//Alternativas del manejo
		//Panic
		//Log

	}
	defer c.Close()

	//Envío gradual y recepción de respuesta
	mensajes := []string{"Hola\n", "Mundo\n", "UniLibre\n", "SistemasDistribuídos\n"}
	for _, mensaje := range mensajes {

		fmt.Println("Enviando Desde el Cliente-> ", mensaje)

		time.Sleep(500 * time.Millisecond)
		//Enviar respuesta
		c.Write([]byte(string(mensaje)))

		//Recibimos la respuesta
		//bd, _ := ioutil.ReadAll(c)
		//Mostramos en pantalla
		//fmt.Println(string(bd))

		//Procesamiento de respuesta versión 2
		respuesta := make([]byte, 1024)
		_, err := c.Read(respuesta)
		if err != nil {
			fmt.Println("Fallo en la lectura de la respuesta del servidor!!! ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Respuesta recibida es = ", string(respuesta))

	}

}
