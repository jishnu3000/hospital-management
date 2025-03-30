package utils

import (
    "errors"
    "github.com/golang-jwt/jwt/v5"
    "time"
    "os"
)

// Claims struct for JWT
type Claims struct {
    UserID string `json:"user_id"`
    jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token for a user
func GenerateToken(userID string) (string, error) {
    jwtSecret := os.Getenv("JWT_SECRET") // Load JWT secret from environment variable
    if jwtSecret == "" {
        return "", errors.New("JWT secret is not set")
    }

    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(jwtSecret)) // Convert to []byte
}

// VerifyToken validates a JWT token and returns user claims
func VerifyToken(tokenString string) (*Claims, error) {
    jwtSecret := os.Getenv("JWT_SECRET") // Load secret key
    if jwtSecret == "" {
        return nil, errors.New("JWT secret is not set")
    }

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte(jwtSecret), nil // Convert to []byte
    })

    if err != nil || !token.Valid {
        return nil, errors.New("invalid token")
    }

    return claims, nil
}
