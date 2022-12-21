package domain

import "time"

type Booking struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	EventID   uint64    `json:"event_id"`
	TicketID  uint      `json:"ticket_id"`
	Amount    uint      `json:"amount"`
	Paid      bool      `json:"paid"`
	CreatedAt time.Time `json:"created_at"`
	User      User
	Event     Event
	Ticket    Ticket
}

type BookingRepository interface{}

type BookingUsecase interface{}
