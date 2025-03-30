package middleware

import (
    "context"
    "hospital-management/utils"
    "net/http"
    "strings"
)

// JWTMiddleware validates JWT tokens
func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization token missing", http.StatusUnauthorized)
            return
        }

        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }

        claims, err := utils.VerifyToken(tokenParts[1])
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Store user ID in request context
        ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
