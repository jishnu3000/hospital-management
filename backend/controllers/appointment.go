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

// Schedule an appointment
func ScheduleAppointment(w http.ResponseWriter, r *http.Request) {
    var appointment models.Appointment
    json.NewDecoder(r.Body).Decode(&appointment)

    appointment.ID = primitive.NewObjectID()
    collection := config.DB.Collection("appointments")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, appointment)
    if err != nil {
        http.Error(w, "Error scheduling appointment", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Appointment scheduled successfully"})
}

// Get all appointments
func GetAppointments(w http.ResponseWriter, r *http.Request) {
    collection := config.DB.Collection("appointments")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error retrieving appointments", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var appointments []models.Appointment
    for cursor.Next(ctx) {
        var appointment models.Appointment
        cursor.Decode(&appointment)
        appointments = append(appointments, appointment)
    }

    json.NewEncoder(w).Encode(appointments)
}

// Cancel an appointment
func CancelAppointment(w http.ResponseWriter, r *http.Request) {
    var appointment models.Appointment
    json.NewDecoder(r.Body).Decode(&appointment)

    collection := config.DB.Collection("appointments")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.DeleteOne(ctx, bson.M{"_id": appointment.ID})
    if err != nil {
        http.Error(w, "Error canceling appointment", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Appointment canceled successfully"})
}
