package domain

import "time"

type Event struct {
	ID                uint64    `json:"id" gorm:"primaryKey"`
	UserID            uint64    `json:"organizer_id" gorm:"not null"`
	Title             string    `json:"title" gorm:"type:varchar(255);not null"`
	EventFormatID     uint      `json:"format_id" gorm:"type:int unsigned;not null"`
	EventTopicID      uint      `json:"topic_id" gorm:"type:int unsigned;not null"`
	Tags              string    `json:"tags"`
	Date              time.Time `json:"date" gorm:"not null"`
	IsOnline          bool      `json:"is_online" gorm:"not null;default:false"`
	IsFullBooked      bool      `json:"is_full_booked" gorm:"not null;default:false"`
	Location          string    `json:"location" gorm:"type:varchar(255);not null"`
	Description       string    `json:"description"`
	TermsAndCondition string    `json:"terms_and_condition"`
	Slug              string    `json:"slug" gorm:"type:varchar(255);not null;unique"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	User              User
	EventFormat       EventFormat
	EventTopic        EventTopic
}

type EventFormat struct {
	ID     uint   `json:"id" gorm:"primarykey;type:int unsigned"`
	Format string `json:"format" gorm:"type:varchar(255);not null;unique"`
}

type EventTopic struct {
	ID    uint   `json:"id" gorm:"primarykey;type:int unsigned"`
	Topic string `json:"topic" gorm:"type:varchar(255);not null;unique"`
}

type EventRepository interface{}

type EventUsecase interface{}
