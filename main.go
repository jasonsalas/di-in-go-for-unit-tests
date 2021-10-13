package main

import (
	"fmt"

	"github.com/jasonsalas/di-in-go-for-unit-tests/astros"
	"github.com/jasonsalas/di-in-go-for-unit-tests/utils"
)

func main() {
	liveClient := astros.LiveGetWebRequest{}
	number, err := utils.GetAstronauts(liveClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("# of astronauts: %d", number)
}
