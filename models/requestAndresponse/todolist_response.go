package requestAndresponse

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TodoListResponse struct {
	TodoID      int       `json:"todo_id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type JwtCustomClaims struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}
