package domain

import "time"

type Ticket struct {
	ID           uint      `json:"id"`
	EventID      uint64    `json:"event_id"`
	TicketTypeID uint      `json:"ticket_type"`
	Name         string    `json:"name"`
	Amount       uint      `json:"amount"`
	Price        float32   `json:"price"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	TicketType   TicketType
	Event        Event
}

type TicketType struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
}

type TicketRepository interface{}

type TicketUsecase interface{}
