package house

import (
	"context"
	"errors"
)

type service struct {
	rep Repository
}

func NewService(repository Repository) Service {
	return &service{
		rep: repository,
	}

}

// CreateHouse handles the creation of a new house
func (s *service) CreateHouse(ctx context.Context, req *CreateHouseRequest) (*CreateHouseResponse, error) {
	// Validate the request
	if req.Name == "" || req.Adress == "" {
		return nil, errors.New("name and address are required")
	}

	// Create House entity from request
	house := &House{
		Name:   req.Name,
		Adress: req.Adress,
	}

	// Use repository to save the house
	createdHouse, err := s.rep.CreateHouse(ctx, house)
	if err != nil {
		return nil, err
	}

	// Create and return the response
	resp := &CreateHouseResponse{
		Name:   createdHouse.Name,
		Adress: createdHouse.Adress,
	}

	return resp, nil
}
