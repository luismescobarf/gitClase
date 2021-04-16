package main

import (
	"fmt"
	"time"
)

func finder(theMine []string) []string {
	var foundOre []string
	for _, m := range theMine {
		if m == "ore1" || m == "ore2" || m == "ore3" {
			foundOre = append(foundOre, m)
		}
	}
	fmt.Println("Encontrados: ", foundOre)
	return foundOre
}

func miner(foundOre []string) []string {
	var minedOre []string
	for _, m := range foundOre {
		minedOre = append(minedOre, m)
	}
	fmt.Println("Extraídos: ", minedOre)
	return minedOre
}

func smelter(minedOre []string) []string {
	var smeltedOre []string
	for _, m := range minedOre {
		smeltedOre = append(smeltedOre, m)
	}
	fmt.Println("Fundidos: ", minedOre)
	return minedOre
}

func main() {

	bufferedChan := make(chan string, 3)

	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()

	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
	}()

	<-time.After(3 * time.Second)

}
