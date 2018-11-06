package main

import (
	"fmt"
	"go-modules/pkg/utils"
)

func main() {
	fmt.Println("Hai", utils.Triple(10))

	err := utils.Fail1("nope")

	fmt.Printf("Error: %+v\n", err)
}
