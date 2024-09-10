// Package user defines the user service and interfaces for handling user operations.
package user

import (
	"context"                                // Package for managing context across API boundaries"
	util "gopractice2/server/database/utils" // Package for utility functions like password hashing
	"strconv"                                // Package for string conversions
	"time"                                   // Package for time manipulation

	"github.com/golang-jwt/jwt/v4" // Package for handling JSON Web Tokens (JWT)
)

// Constant for the JWT secret key
const (
	secretKey = "secret"
)

// service struct implements the Service interface using a Repository and a timeout duration.
type service struct {
	Repository // Embedded Repository interface for database operation
	// redisClient *redis.Client
	timeout time.Duration // Timeout duration for context
}

// NewService returns a new instance of service that implements the Service interface.
func NewService(repository Repository) Service {
	return &service{
		repository,
		// rd,
		time.Duration(2) * time.Second, // Sets the context timeout to 2 seconds

	}
}

// CreateUser handles the creation of a new user.
// It takes a context and a CreateUserReq struct as parameters and returns a CreateUserRes struct.
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	// Create a context with a timeout and ensure it gets cancelled to free resources
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Hash the user's password using a utility function
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err // Return error if password hashing fails
	}

	// Create a new User struct with the hashed password
	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Save the new user to the database using the repository
	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err // Return error if database operation fails
	}

	// Create a response struct with the user details
	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil // Return the created user response
}

// MyJWTClaims defines custom claims for the JWT, including user ID and username.
type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Login handles user login and JWT generation.
// It takes a context and a LoginUserReq struct as parameters and returns a LoginUserRes struct.
func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	// Create a context with a timeout and ensure it gets cancelled to free resources
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	//check if the user is cached in redis
	// catchedData, err := s.redisClient.Get(ctx, req.Email).Result()

	// if err == nil {
	// 	var cachedUser User
	// 	if json.Unmarshal([]byte(catchedData), &cachedUser) == nil {
	// 		err =  util.CheckPassword(req.Password, cachedUser.Password)
	// 		if( err != nil ) {
	// 			return nil, err
	// 		}
	// 		token := jwt.NewWithClaims(jwt.SigningMethodES256, MyJWTClaims{
	// 			ID: strconv.Itoa(int(cachedUser.ID)),
	// 			Username: cachedUser.Username,
	// 			RegisteredClaims: jwt.RegisteredClaims{
	// 				Issuer: strconv.Itoa(int(cachedUser.ID)),
	// 				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24*time.Hour)),
	// 			},
	// 		})
	// 		ss, err := token.SignedString([]byte(secretKey))
	// 		if err != nil {
	// 			return &LoginUserRes{}, err
	// 		}
	// 		return &LoginUserRes{
	// 			accessToken: ss ,
	// 			Username : cachedUser.Username,
	// 			ID: strconv.Itoa(int(cachedUser.ID)),
	// 		}, nil
	// 	}

	// }

	// Retrieve the user by email from the repository
	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err // Return error if user retrieval fails
	}

	// Check if the provided password matches the stored hashed password
	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err // Return error if password check fails
	}

	// Create JWT claims with user ID and username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Set token expiry to 24 hours
		},
	})

	// Sign the JWT with the secret key
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err // Return error if JWT signing fails
	}

	// Create a response struct with the JWT and user details
	return &LoginUserRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

// Example function to add user ID to online users Set in Redis
// func (h *service) AddOnlineUser(userID int64) error {
// 	return h.redisClient.SAdd(context.Background(), "online_users", userID).Err()
// }

// // Example function to remove user ID from online users Set in Redis
// func (h *service) RemoveOnlineUser(userID int64) error {
// 	return h.redisClient.SRem(context.Background(), "online_users", userID).Err()
// }

// // Example function to get all online user IDs from Redis
// func (h *service) GetOnlineUsers() ([]string, error) {
// 	return h.redisClient.SMembers(context.Background(), "online_users").Result()
// }
