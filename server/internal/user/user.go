package user

// CREATE TABLE User (
//     user_id INT AUTO_INCREMENT PRIMARY KEY,
//     house_id INT NOT NULL,
//     room_id INT,
//     username VARCHAR(50) NOT NULL UNIQUE,
//     email VARCHAR(50) NOT NULL UNIQUE,
//     password VARCHAR(255) NOT NULL,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     FOREIGN KEY (house_id) REFERENCES House(house_id),
//     FOREIGN KEY (room_id) REFERENCES Room(room_id)
// );

import "context"

// User struct represents a user in the system.
type User struct {
	ID       int64  `json:"id"`       // Unique identifier for the user
	House_id int64  `json:"house_id"` // Unique identifier for the
	Room_id  int64  `json:"room_id"`  // Unique identifier for the
	Username string `json:"username"` // Username of the user
	Email    string `json:"email"`    // Email of the user
	Password string `json:"password"` // Hashed password of the user
}

// CreateUserReq struct represents the payload for creating a new user.
type CreateUserReq struct {
	Username string `json:"username"` // Username for the new user
	Email    string `json:"email"`    // Email for the new user
	Password string `json:"password"` // Password for the new user
}

// CreateUserRes struct represents the response after creating a new user.
type CreateUserRes struct {
	ID       string `json:"id"`       // Unique identifier for the new user
	Username string `json:"username"` // Username of the new user
	Email    string `json:"email"`    // Email of the new user
}

// LoginUserReq struct represents the payload for logging in a user.
type LoginUserReq struct {
	Email    string `json:"email"`    // Email of the user trying to log in
	Password string `json:"password"` // Password of the user trying to log in
}

// LoginUserRes struct represents the response after logging in a user.
type LoginUserRes struct {
	accessToken string // JWT access token for the logged-in user
	ID          string `json:"id"`       // Unique identifier for the logged-in user
	Username    string `json:"username"` // Username of the logged-in user
}

type UserRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// Repository interface defines the methods for interacting with the user data store.
type Repository interface {
	// CreateUser inserts a new user into the data store and returns the created user.
	CreateUser(ctx context.Context, user *User) (*User, error)
	// GetUserByEmail retrieves a user by their email from the data store.
	GetUserByEmail(ctx context.Context, email string) (*User, error)

	GetUserById(ctx context.Context, id int) (*User, error)
}

// Service interface defines the methods for handling user-related business logic.
type Service interface {
	// CreateUser creates a new user and returns the response structure.
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	// Login authenticates a user and returns the response structure with an access token.
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)

	GetUserById(c context.Context, id int) (*UserRes, error)
	// Add online user into redis
	// 	AddOnlineUser(userID int64) error
	// 	// Remove online user from redis
	// 	RemoveOnlineUser(userID int64) error
	// 	// Get all the users from redis set
	// GetOnlineUsers() ([]string, error) //return an array of strings or error

	//
}
