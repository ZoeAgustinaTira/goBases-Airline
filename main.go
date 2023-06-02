package main

import (
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets"
	"github.com/ZoeAgustinaTira/goBases-Airline/reader"
)

func main() {
	ch := make(chan string)

	reader.ReadFile()

	go tickets.AverageDestination(ch, "Brazil")

	total, err := tickets.GetTotalTicketsByDestination("Brazil")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(total)

	go tickets.GetCountByPeriod(ch, "night")

	fmt.Println(<-ch)

}
