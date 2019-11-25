package vo

import "gojstest/model"

// TicketStatus is Tickets' status.
type TicketStatus int

const (
	TicketStatusNew TicketStatus = iota
	TicketStatusPaid
	TicketStatusCleared
	TicketStatusCancelled
	TicketStatusDeleted
)

// Validate validates TicketStatus value.
func (ts TicketStatus) Validate() error {
	// allcases
	switch ts {
	case TicketStatusNew, TicketStatusPaid, TicketStatusCleared, TicketStatusCancelled, TicketStatusDeleted:
		return nil
	}

	return model.ErrInvalidF("invalid TicketStatus:%v", ts)
}

// Available returns true if it's available.
func (ts TicketStatus) Available() bool {
	// allcases
	switch ts {
	case TicketStatusNew, TicketStatusPaid, TicketStatusCleared:
		return true
	case TicketStatusCancelled, TicketStatusDeleted:
		return false
	}

	panic("unreacheable")
}
