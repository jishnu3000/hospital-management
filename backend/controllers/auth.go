package controllers

import (
    "encoding/json"
    "hospital-management/config"
    "hospital-management/models"
    "hospital-management/utils"
    "net/http"
    "context"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

// Register a new user
// RegisterUser handles new user registration
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Hash the password before storing
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Failed to hash password", http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    // Insert user into MongoDB
    collection := config.DB.Collection("users")
    _, err = collection.InsertOne(context.TODO(), user)
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUser handles user authentication
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var input models.User
    var storedUser models.User

    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
        return
    }

    // Check if user exists in MongoDB
    collection := config.DB.Collection("users")
    err = collection.FindOne(context.TODO(), bson.M{"email": input.Email}).Decode(&storedUser)
    w.Header().Set("Content-Type", "application/json") // Set header here
    if err == mongo.ErrNoDocuments {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"error": "Email not found"})
        return
    } else if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Database error"})
        return
    }

    // Compare passwords
    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(input.Password))
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]string{"error": "Incorrect password"})
        return
    }

    // Generate JWT
    token, err := utils.GenerateToken(storedUser.ID.Hex())
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(map[string]string{"error": "Failed to generate token"})
        return
    }

    // Send success response
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}