package main

import (
	"github.com/ZoeAgustinaTira/goBases-Airline/internal/tickets"
	"github.com/ZoeAgustinaTira/goBases-Airline/reader"
)

func main() {
	reader.ReadFile()
	//total, err := tickets.GetTotalTickets("Brazil")

	tickets.GetCountByPeriod("noche")

}
