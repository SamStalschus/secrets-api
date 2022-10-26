package internal

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User of account
type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name      string             `json:"name,omitempty" bson:"name"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updatedAt"`
}

// AuthUser represent data of user authenticate
type AuthUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token represent data token to return
type Token struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
