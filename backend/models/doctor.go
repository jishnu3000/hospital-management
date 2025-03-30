package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct {
    ID              primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
    Name            string               `bson:"name" json:"name"`
    Specialization  string               `bson:"specialization" json:"specialization"`
    Contact         string               `bson:"contact" json:"contact"`
    AssignedPatients []primitive.ObjectID `bson:"assigned_patients,omitempty" json:"assigned_patients"`
}
