package main

import (
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets"
	"github.com/ZoeAgustinaTira/goBases-Airline/reader"
)

func main() {
	reader.ReadFile()

	total, err := tickets.GetTotalTicketsByDestination("Brazil")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(total)

	ch := make(chan string)

	go tickets.GetCountByPeriod(ch, "night")

	fmt.Println(<-ch)


}
