package main

import (
	"fmt"

	"github.com/jasonsalas/isolate-astros/astros"
	"github.com/jasonsalas/isolate-astros/utils"
)

func main() {
	liveClient := astros.LiveGetWebRequest{}
	number, err := utils.GetAstronauts(liveClient)
	if err != nil {
		panic(err)
	}
	fmt.Printf("# of astronauts: %d", number)
}
