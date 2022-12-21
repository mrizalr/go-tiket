package main

import (
	"github.com/mrizalr/go-tiket/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:2252)/tiketdb"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&domain.User{},
		&domain.Credentials{},
		&domain.Event{},
		&domain.EventFormat{},
		&domain.EventTopic{},
		&domain.Ticket{},
		&domain.TicketType{},
		&domain.Booking{},
	)
}
