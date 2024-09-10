package message

import (
	"context"
	"time"
)

type Message struct {
	Content   string    `json:"content"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	TimeStamp time.Time `json:"time_stamp"`
}

// type Repository interface {
// 	CreateHouse(ctx context.Context, house *House) (*House, error)
// }

//	type Service interface {
//		CreateHouse(c context.Context, req *CreateHouseRequest) (*CreateHouseResponse, error)
//	}
//
// dependency injection plug
type Repository interface {
	SaveMessage(ctx context.Context, message *Message) error
	GetMessageHistory(ctx context.Context, sender string, receiver string) ([]Message, error)
}
