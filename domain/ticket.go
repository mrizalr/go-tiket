package domain

import "time"

type Ticket struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	EventID      uint64    `json:"event_id" gorm:"not null"`
	TicketTypeID uint      `json:"ticket_type" gorm:"not null"`
	Name         string    `json:"name" gorm:"type:varchar(255); not null"`
	Amount       uint      `json:"amount" gorm:"type:int unsigned;not null;default:0"`
	Price        float32   `json:"price" gorm:"not null;default:0"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date" gorm:"not null"`
	EndDate      time.Time `json:"end_date" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TicketType   TicketType
	Event        Event
}

type TicketType struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Type string `json:"type" gorm:"type:varchar(255);not null;unique"`
}

type TicketRepository interface{}

type TicketUsecase interface{}
