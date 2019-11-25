package main

import (
	"gojstest/model/entity"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("ticketFromJSON", func(s string) (*entity.Ticket, error) {
		return entity.TicketFromJSON(s)
	})
}
