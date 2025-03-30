package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name        string             `bson:"name" json:"name"`
    Age         int                `bson:"age" json:"age"`
    Gender      string             `bson:"gender" json:"gender"`
    Contact     string             `bson:"contact" json:"contact"`
    Address     string             `bson:"address" json:"address"`
    MedicalHistory []string        `bson:"medical_history" json:"medical_history"`
}
