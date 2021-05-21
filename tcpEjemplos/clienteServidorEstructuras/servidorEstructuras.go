package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

type Mensaje struct {
	Id       string
	Mensaje  string
	Promedio float32
}

func manejarConexiónBidireccional(c net.Conn) {

	//Pendiente cierre de conexión
	defer c.Close()

	fmt.Printf("Sirviendo %s \n", c.RemoteAddr().String())

	for {

		//Contenedor temporal
		tmp := make([]byte, 1024)

		_, err := c.Read(tmp)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Convertir bytes en un buffer
		tmpbuff := bytes.NewBuffer(tmp)

		//Instancia de la estructura
		tmpstruct := new(Mensaje)

		//Instanciar el decodificador del objeto (serializado)
		gobobj := gob.NewDecoder(tmpbuff)

		//Decodificar buffer en la estructura (unmarshals)
		gobobj.Decode(tmpstruct)

		//Mostrar lo que se recibió del cliente
		fmt.Println("Recibido: ", tmpstruct)

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
