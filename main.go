package main

import (
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets"
	"github.com/ZoeAgustinaTira/goBases-Airline/reader"
)

func main() {
	reader.ReadFile()
	avg, err := tickets.AverageDestination("Brazil")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(avg)
}
