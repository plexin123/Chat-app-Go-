package house

import "context"

type House struct {
	ID     int64  `json:"id"`
	Name   string `json:"username"`
	Adress string `json:"address"`
	// NumberofStudents string `json:"numberofstundets"`
}

type CreateHouseRequest struct {
	Name   string `json:"name"`
	Adress string `json:"address"`
	// NumberofStudents string `json:"numberofstundets"`

}
type CreateHouseResponse struct {
	Name   string `json:"name"`
	Adress string `json:"address"`
	// NumberofStudents string `json:"numberofstundets"`
}

type Repository interface {
	CreateHouse(ctx context.Context, house *House) (*House, error)
}

type Service interface {
	CreateHouse(c context.Context, req *CreateHouseRequest) (*CreateHouseResponse, error)
}
