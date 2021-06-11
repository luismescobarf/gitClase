package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"sync"
)

var contenedorMapeo [][]tupla

type tupla struct {
	llave string
	valor int
}

type Mensaje struct {
	Mensaje string
}

func manejarConexiónBidireccional(c net.Conn) {

	var wg sync.WaitGroup
	//Pendiente cierre de conexión
	defer c.Close()

	//fmt.Printf("Sirviendo %s \n", c.RemoteAddr().String())

	for {

		inicio := time.Now()
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

		lineas, _ := leerLineas(tmpstruct.Mensaje)

		acumuladorPartesTrabajo := organizarLineas(lineas)

		//Definición de función para aplicar en el map
		//Función simple que no depende de otros resultados
		//y puede aplicarse Concurrente/Paralelamente
		funcion := func(palabra string) tupla {
			return tupla{llave: palabra, valor: 1}
		}

		for _, parteTrabajo := range acumuladorPartesTrabajo {
			wg.Add(1)
			go dividirSlice(parteTrabajo,funcion, &wg)
		}
		//fmt.Println("Muy buenas  a todos")
		
		wg.Wait()
		//fmt.Println(contenedorMapeo)
		var contenedorShuffle map[string][]tupla
		contenedorShuffle = shuffler(contenedorMapeo)

		var contenedorReducer map[string]int
		contenedorReducer = reducer(contenedorShuffle)

		fmt.Println("********>Resultado final")
		for llave, valor := range contenedorReducer {
			fmt.Printf("%s, %d\n", llave, valor)
		}

		fmt.Println("---------------------------------------")
		fmt.Println("Número de palabras: ", len(contenedorReducer))
		fmt.Println("Tiempo total de cómputo:", time.Since(inicio))
	}
}

//Funcion para dividir en slice
func dividirSlice(parteTrabajo []string , funcion func(string) tupla,wg *sync.WaitGroup){
	defer wg.Done()
	contenedorMapeo = append(contenedorMapeo, etapaMap(parteTrabajo, funcion))
	//fmt.Println(contenedorMapeo)
}

//Función para lectura de un archivo
func leerLineas(path string) ([]string, error) {
	archivo, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	//Arreglo que recibirá las líneas del archivo
	var lineas []string
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		lineas = append(lineas, scanner.Text())
	}
	//Retornar las líneas y el estado de la lectura
	return lineas, scanner.Err()
}

func organizarLineas(lineas []string) [][]string {
	var acumuladorPartesTrabajo [][]string
	for _, linea := range lineas {
		//Limpiar si hay caracteres de movimientos de carro al comienzo y al final
		linea = strings.TrimSpace(linea)
		//Quitar puntuaciones
		linea = strings.ReplaceAll(linea, ":", "")
		linea = strings.ReplaceAll(linea, ";", "")
		linea = strings.ReplaceAll(linea, ",", "")
		linea = strings.ReplaceAll(linea, ".", "")
		linea = strings.ReplaceAll(linea, "!", "")
		linea = strings.ReplaceAll(linea, "?", "")
		linea = strings.ReplaceAll(linea, "¿", "")
		linea = strings.ReplaceAll(linea, "«", "")
		linea = strings.ReplaceAll(linea, "»", "")
		linea = strings.ReplaceAll(linea, "(", "")
		linea = strings.ReplaceAll(linea, ")", "")
		//linea = strings.ReplaceAll(linea, "\"", "")
		//linea = strings.ReplaceAll(linea, "\'", "")
		linea = strings.ReplaceAll(linea, "-", "")
		//Separar las palabras de la línea o string recibido
		listadoPalabras := strings.Split(linea, " ")
		//fmt.Println(listadoPalabras)
		acumuladorPartesTrabajo = append(acumuladorPartesTrabajo, listadoPalabras)
		//fmt.Println(acumuladorPartesTrabajo)
	}
	//fmt.Println(acumuladorPartesTrabajo)
	return acumuladorPartesTrabajo
}

//Etapa de mapeo para la porción de trabajo recibida
func etapaMap(listadoPalabras []string, funcion func(string) tupla) []tupla {
	//Contenedor con el resultado del mapeo
	contenedorLlaveValor := make([]tupla, len(listadoPalabras))
	//Realizar el mapeo
	for i, palabra := range listadoPalabras {
		contenedorLlaveValor[i] = funcion(palabra)
	}
	//Retornar el resultado del mapeo (trabajo del mapper)
	//fmt.Println(contenedorLlaveValor)
	return contenedorLlaveValor
}

//Etapa de "barajado" del mapeo
func shuffler(contenedorLlaveValor [][]tupla) map[string][]tupla {
	//Preprar el contenedor de salida del barajado
	apilados := make(map[string][]tupla)
	//Por cada una de las líneas o pedazos resultantes de mapeo
	for _, linea := range contenedorLlaveValor {
		//fmt.Println(linea)
		//Apilar el resultado del mapeo por llaves
		for _, kv := range linea {
			apilados[kv.llave] = append(apilados[kv.llave], kv)
			//fmt.Println(apilados[kv.llave])
		}
	}
	//Retornar el resultado de apilar comunes
	return apilados
}

//Etapa reduce
func reducer(contenedorShuffler map[string][]tupla) map[string]int {
	//Preparar el contenedor de reducción
	contenedorReducer := make(map[string]int)
	//Reducir cada una de las pilas o montones
	for _, monton := range contenedorShuffler {
		for _, kv := range monton {
			contenedorReducer[kv.llave] += kv.valor
		}
	}
	//Retornar el resultado de la reducción
	return contenedorReducer
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
