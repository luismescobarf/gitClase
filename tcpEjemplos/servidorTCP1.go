package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func manejarConexiónUnidireccional(c net.Conn) {

	fmt.Printf("Sirviendo %s \n", c.RemoteAddr().String())

	datosEntrantes, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(datosEntrantes)

	c.Close()

}

func manejarConexiónBidireccional(c net.Conn) {

	//Pendiente cierre de conexión
	defer c.Close()

	fmt.Printf("Sirviendo %s \n", c.RemoteAddr().String())

	for {

		datosEntrantes, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		//fmt.Println(datosEntrantes)

		//Criterio de parada para conexión
		temp := strings.TrimSpace(string(datosEntrantes))
		if temp == "DETENERSE!" {
			break
		}

		//Procesar la respuesta con base en el valor ingresado
		respuesta := strings.ToUpper(temp)
		respuesta = respuesta + "\n"

		//Enviar respuesta
		c.Write([]byte(string(respuesta)))

		//Mostrar lo que se le envió al cliente
		fmt.Println("****Enviado: ", respuesta)

	}

}

func main() {

	var Puerto string
	argumentos := os.Args
	if len(argumentos) == 1 {
		Puerto = ":12345"
	} else {
		Puerto = ":" + argumentos[1]
	}
	l, err := net.Listen("tcp4", Puerto)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("Demonio listo!")

	//Demonio -> Escucha permanente sobre el puerto establecido
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		//Está establecida la conexión y se procede a servir la conexión establecida
		//manejarConexiónUnidireccional(c)
		go manejarConexiónBidireccional(c)

	}

}
