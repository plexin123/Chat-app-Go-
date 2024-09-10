package room

import (
	"context"
)

type service struct {
	rep Repository
}

func NewService(rep Repository) Service {
	return &service{
		rep: rep,
	}
}

// type Room struct {
// 	ID         int64     `json:"id"`
// 	Name       string    `json:"name"`
// 	House_id   int64     `json:"house_id"`
// 	User_id    int64     `json:"user_id"`
// 	Status     bool      `json:"status"`
// 	Created_at time.Time `json:"created_at"`
// }

func (s *service) CreateRoom(c context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error) {

	reqs := &Room{
		Name:     req.Name,
		House_id: req.House_id,
		User_id:  req.User_id,
		Status:   req.Status,
	}

	r, err := s.rep.CreateRoom(c, reqs)
	if err != nil {
		return nil, err
	}

	res := &CreateRoomResponse{

		// assign the the response to res
		ID:       r.ID,
		Name:     r.Name,
		House_id: r.House_id,
		User_id:  r.User_id,
		Status:   r.Status,
	}
	return res, nil

}
