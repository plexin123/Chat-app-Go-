package room

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serv Service
}

func NewHandler(serv Service) *handler {
	return &handler{
		serv: serv,
	}

}

// func (s *handler) CreateHouse(c *gin.Context) {
// 	var req CreateHouseRequest
// 	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
// 		return
// 	}

// 	res, err := s.serv.CreateHouse(c.Request.Context(), &req)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, res)

// }

func (h *handler) CreateRoom(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	res, err := h.serv.CreateRoom(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, res)
}
