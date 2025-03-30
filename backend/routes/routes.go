package routes

import (
    "hospital-management/controllers"
    "hospital-management/middleware"

    "github.com/gorilla/mux"
)

// InitializeRoutes sets up the API routes
func InitializeRoutes() *mux.Router {
    router := mux.NewRouter()

    // Authentication Routes
    router.HandleFunc("/api/register", controllers.RegisterUser).Methods("POST")
    router.HandleFunc("/api/login", controllers.LoginUser).Methods("POST")

    // Protected Routes (Require JWT Authentication)
    protectedRoutes := router.PathPrefix("/api").Subrouter()
    protectedRoutes.Use(middleware.JWTMiddleware)

    // Patient Routes
    protectedRoutes.HandleFunc("/patients", controllers.AddPatient).Methods("POST")
    protectedRoutes.HandleFunc("/patients", controllers.GetPatients).Methods("GET")
    protectedRoutes.HandleFunc("/patients", controllers.UpdatePatient).Methods("PUT")
    protectedRoutes.HandleFunc("/patients", controllers.DeletePatient).Methods("DELETE")

    // Doctor Routes
    protectedRoutes.HandleFunc("/doctors", controllers.AddDoctor).Methods("POST")
    protectedRoutes.HandleFunc("/doctors", controllers.GetDoctors).Methods("GET")
    protectedRoutes.HandleFunc("/doctors/assign", controllers.AssignPatientToDoctor).Methods("POST")

    // Appointment Routes
    protectedRoutes.HandleFunc("/appointments", controllers.ScheduleAppointment).Methods("POST")
    protectedRoutes.HandleFunc("/appointments", controllers.GetAppointments).Methods("GET")
    protectedRoutes.HandleFunc("/appointments", controllers.CancelAppointment).Methods("DELETE")

    return router
}
