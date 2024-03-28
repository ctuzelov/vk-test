package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ad struct {
	Id          primitive.ObjectID `bson:"_id"`
	ID          int                `bson:"id" json:"id"`
	ProjectName string             `bson:"project_name,omitempty" json:"project_name,omitempty"`
	Category    string             `bson:"category,omitempty" json:"category,omitempty"`
	ProjectType string             `bson:"project_type,omitempty" json:"project_type,omitempty"`
	AgeCategory string             `bson:"age_category,omitempty" json:"age_category,omitempty"`
	Year        int                `bson:"year,omitempty" json:"year,omitempty"`
	RunningTime int                `bson:"running_time,omitempty" json:"running_time,omitempty"`
	Keywords    string             `bson:"keywords,omitempty" json:"keywords,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	Director    string             `bson:"director,omitempty" json:"director,omitempty"`
	Producer    string             `bson:"producer,omitempty" json:"producer,omitempty"`
}
