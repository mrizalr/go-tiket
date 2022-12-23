package domain

import "time"

type Booking struct {
	ID        uint64    `json:"id" gorm:"primaryKey"`
	UserID    uint64    `json:"user_id" gorm:"not null"`
	EventID   uint64    `json:"event_id" gorm:"not null"`
	TicketID  uint      `json:"ticket_id" gorm:"not null"`
	Amount    uint      `json:"amount" gorm:"type:int unsigned;not null"`
	Paid      bool      `json:"paid" gorm:"not null;default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	User      User
	Event     Event
	Ticket    Ticket
}

type BookingRepository interface{}

type BookingUsecase interface{}
