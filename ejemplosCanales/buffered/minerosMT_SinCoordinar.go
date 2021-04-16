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
	theMine := []string{"rock", "ore1", "ore2", "rock", "ore3"}
	//smelter(miner(finder(theMine)))
	go finder(theMine)
	go finder(theMine)

	<-time.After(3 * time.Second)

}
