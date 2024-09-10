package user

// Import necessary packages
import (
	"net/http" // Package for HTTP client and server implementations

	"github.com/gin-gonic/gin" // Package for Gin, a web framework
	// "github.com/redis/go-redis/v9"
)

// Handler struct embeds the Service interface to provide the necessary methods
type Handler struct {
	Service // This field should be an interface or struct that provides service methods
	// redisClient *redis.Client // create redis client to manage online users and storing user metadata
}

// NewHandler initializes a new Handler with the given service
func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
		// redisClient: rc,
	}
}

// CreateUser handles the user creation endpoint
func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq // Struct for creating a user request payload
	// Bind JSON payload to CreateUserReq struct
	if err := c.ShouldBindJSON(&u); err != nil {
		// Return HTTP 400 status code with error message if binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the CreateUser method from the service with the request context and payload
	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		// Return HTTP 500 status code with error message if service call fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return HTTP 200 status code with the response from the service
	c.JSON(http.StatusOK, res)
}

// Login handles the user login endpoint
func (h *Handler) Login(c *gin.Context) {
	var user LoginUserReq // Struct for login user request payload
	// Bind JSON payload to LoginUserReq struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// Return HTTP 400 status code with error message if binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the Login method from the service with the request context and payload
	u, err := h.Service.Login(c.Request.Context(), &user)
	if err != nil {
		// Return HTTP 500 status code with error message if service call fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// userId, _ := strconv.ParseInt(u.ID, 10, 64)
	// if err := h.Service.AddOnlineUser(userId); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to online users"})
	// 	return
	// }

	// Set a JWT cookie with the access token from the service response
	c.SetCookie("jwt", u.accessToken, 60*60*24, "/", "localhost", false, true)
	// Return HTTP 200 status code with the response from the service
	c.JSON(http.StatusOK, u)
}

// Logout handles the user logout endpoint
func (h *Handler) Logout(c *gin.Context) {
	// Invalidate the JWT cookie by setting its expiration to a past date
	c.SetCookie("jwt", "", -1, "", "", false, true)
	// Return HTTP 200 status code with a logout successful message
	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

// func (h *Handler) getUsers(c *gin.Context) {
// 	users, err := h.GetOnlineUsers()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch all the users"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, users)
// }
//apply redis caching
