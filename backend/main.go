package main

import (
    "hospital-management/config"
    "hospital-management/middleware"
    "hospital-management/routes"
    "log"
    "net/http"
)

func main() {
    config.LoadEnv()
    config.ConnectDB() // Connect to MongoDB

    router := routes.InitializeRoutes()

    // Apply CORS Middleware
    handler := middleware.CORSMiddleware(router)

    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
