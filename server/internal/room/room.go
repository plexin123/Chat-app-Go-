package room

import (
	"context"
	"time"
)

type Room struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	House_id   int64     `json:"house_id"`
	User_id    int64     `json:"user_id"`
	Status     bool      `json:"status"`
	Created_at time.Time `json:"created_at"`
}

type CreateRoomRequest struct {
	Name     string `json:"name"`
	House_id int64  `json:"house_id"`
	User_id  int64  `json:"user_id"`
	Status   bool   `json:"status"`
}

type CreateRoomResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	House_id int64  `json:"house_id"`
	User_id  int64  `json:"user_id"`
	Status   bool   `json:"status"`
}

type Repository interface {
	CreateRoom(ctx context.Context, room *Room) (*Room, error)
}

type Service interface {
	CreateRoom(c context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error)
}
