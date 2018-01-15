package main

import (
	"log"

	"github.com/emman27/aoc2017/passphrase"
	"github.com/emman27/aoc2017/spiralMemory"
)

func main() {
	log.Println(spiralMemory.PartA(325489))
	log.Println(spiralMemory.PartB(325489))
	log.Println(passphrase.PartA(passphrase.ActualData))
}
