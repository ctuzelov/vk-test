package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty" validate:"required,email"`
	UserType      string             `bson:"user_type,omitempty" json:"user_type,omitempty" validate:"required"`
	Password      string             `bson:"password,omitempty" json:"password,omitempty" validate:"required,min=6"`
	Token         string             `bson:"token,omitempty" json:"token,omitempty"`
	Refresh_Token string             `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	User_id       string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
}
