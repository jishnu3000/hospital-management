package controllers

import (
    "context"
    "encoding/json"
    "hospital-management/config"
    "hospital-management/models"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

// Add a new doctor
func AddDoctor(w http.ResponseWriter, r *http.Request) {
    var doctor models.Doctor
    json.NewDecoder(r.Body).Decode(&doctor)

    doctor.ID = primitive.NewObjectID()
    collection := config.DB.Collection("doctors")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, doctor)
    if err != nil {
        http.Error(w, "Error adding doctor", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Doctor added successfully"})
}

// Get all doctors
func GetDoctors(w http.ResponseWriter, r *http.Request) {
    collection := config.DB.Collection("doctors")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error retrieving doctors", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var doctors []models.Doctor
    for cursor.Next(ctx) {
        var doctor models.Doctor
        cursor.Decode(&doctor)
        doctors = append(doctors, doctor)
    }

    json.NewEncoder(w).Encode(doctors)
}

// Assign a patient to a doctor
func AssignPatientToDoctor(w http.ResponseWriter, r *http.Request) {
    var assignment struct {
        DoctorID primitive.ObjectID `json:"doctor_id"`
        PatientID primitive.ObjectID `json:"patient_id"`
    }
    json.NewDecoder(r.Body).Decode(&assignment)

    collection := config.DB.Collection("doctors")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    update := bson.M{"$push": bson.M{"assigned_patients": assignment.PatientID}}
    _, err := collection.UpdateOne(ctx, bson.M{"_id": assignment.DoctorID}, update)
    if err != nil {
        http.Error(w, "Error assigning patient", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Patient assigned to doctor successfully"})
}
