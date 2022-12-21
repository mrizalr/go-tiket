package domain

import "time"

type Event struct {
	ID                uint64    `json:"id"`
	UserID            uint64    `json:"organizer_id"`
	Title             string    `json:"title"`
	EventFormatID     uint      `json:"format_id"`
	EventTopicID      uint      `json:"topic_id"`
	Tags              string    `json:"tags"`
	Date              time.Time `json:"date"`
	IsOnline          bool      `json:"is_online"`
	IsFullBooked      bool      `json:"is_full_booked"`
	Location          string    `json:"location"`
	Description       string    `json:"description"`
	TermsAndCondition string    `json:"terms_and_condition"`
	Slug              string    `json:"slug"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	User              User
	EventFormat       EventFormat
	EventTopic        EventTopic
}

type EventFormat struct {
	ID     uint   `json:"id"`
	Format string `json:"format"`
}

type EventTopic struct {
	ID    uint   `json:"id"`
	Topic string `json:"topic"`
}

type EventRepository interface{}

type EventUsecase interface{}
