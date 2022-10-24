package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User of account
type User struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Name      string             `json:"name,omitempty"`
	Email     string             `json:"email,omitempty"`
	Password  string             `json:"password,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty"`
}
