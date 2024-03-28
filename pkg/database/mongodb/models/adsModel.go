package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ad struct {
	ID           primitive.ObjectID `bson:"_id"`
	Title        string             `bson:"title" json:"title"`
	Text         string             `bson:"text" json:"text"`
	ImageURL     string             `bson:"image_url" json:"image_url"`
	Price        float64            `bson:"price" json:"price"`
	CreaterEmail string             `bson:"creater_email" json:"createrEmail"`
	CreateAt     time.Time          `bson:"create_ad" json:"createAd"`
	Owned        bool               `bson:"owned" json:"owned"`
}
