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

// Add a new patient
func AddPatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    json.NewDecoder(r.Body).Decode(&patient)

    patient.ID = primitive.NewObjectID()
    collection := config.DB.Collection("patients")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, patient)
    if err != nil {
        http.Error(w, "Error adding patient", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Patient added successfully"})
}

// Get all patients
func GetPatients(w http.ResponseWriter, r *http.Request) {
    collection := config.DB.Collection("patients")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        http.Error(w, "Error retrieving patients", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var patients []models.Patient
    for cursor.Next(ctx) {
        var patient models.Patient
        cursor.Decode(&patient)
        patients = append(patients, patient)
    }

    json.NewEncoder(w).Encode(patients)
}

// Update a patient
func UpdatePatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    json.NewDecoder(r.Body).Decode(&patient)

    collection := config.DB.Collection("patients")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    filter := bson.M{"_id": patient.ID}
    update := bson.M{"$set": patient}

    _, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        http.Error(w, "Error updating patient", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Patient updated successfully"})
}

// Delete a patient
func DeletePatient(w http.ResponseWriter, r *http.Request) {
    var patient models.Patient
    json.NewDecoder(r.Body).Decode(&patient)

    collection := config.DB.Collection("patients")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := collection.DeleteOne(ctx, bson.M{"_id": patient.ID})
    if err != nil {
        http.Error(w, "Error deleting patient", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Patient deleted successfully"})
}
