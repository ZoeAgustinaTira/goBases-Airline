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

	go tickets.GetTotalTicketsByDestination(ch, "Brazil")

	go tickets.GetCountByPeriod(ch, "night")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
