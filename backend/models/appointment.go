package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Appointment struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    PatientID primitive.ObjectID `bson:"patient_id" json:"patient_id"`
    DoctorID  primitive.ObjectID `bson:"doctor_id" json:"doctor_id"`
    DateTime  time.Time          `bson:"date_time" json:"date_time"`
    Status    string             `bson:"status" json:"status"`
}
