package entity

import (
	"encoding/json"
	"gojstest/model/vo"
	"time"
)

// Ticket describes travel ticket
type Ticket struct {
	Status         vo.TicketStatus `json:"status"`         // status
	ExpirationTime time.Time       `json:"expirationTime"` // expiration time
}

// Available returns true if it's available.
func (t *Ticket) Available(tm time.Time) bool {
	return t.Status.Available() && tm.Before(t.ExpirationTime)
}

// TicketFromJSON unmarshals Ticket from JSON string.
func TicketFromJSON(s string) (*Ticket, error) {
	t := &Ticket{}
	if err := json.Unmarshal([]byte(s), t); err != nil {
		return nil, err
	}
	return t, nil
}
