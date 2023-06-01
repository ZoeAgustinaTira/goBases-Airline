package main

import (
	//	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets"
	"fmt"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/domain"
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/reader"
)

func main() {
	reader.ReadFile()
	fmt.Println(domain.TicketList)

	//total, err := tickets.GetTotalTickets("Brazil")
}
