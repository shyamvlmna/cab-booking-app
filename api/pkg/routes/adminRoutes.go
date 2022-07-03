package routes

import (
	"github.com/gorilla/mux"

	"github.com/shayamvlmna/cab-booking-app/pkg/controllers"
)

func AdminRoutes(r *mux.Router) {

	r.HandleFunc("/admin", controllers.AdminIndex)
	adminRouter := r.PathPrefix("/admin").Subrouter()
	// adminRouter.HandleFunc("/", controllers.AdminIndex)

	adminRouter.HandleFunc("/create", controllers.CreateAdmin).Methods("POST")

	adminRouter.HandleFunc("/login", controllers.AdminLogin)
	adminRouter.HandleFunc("/managedrivers", controllers.ManageDrivers)
	adminRouter.HandleFunc("/login", controllers.AdminLogin)
	adminRouter.HandleFunc("/approve", controllers.ApproveDriver).Methods("POST")

	adminRouter.HandleFunc("/manageusers", controllers.ManageUsers).Methods("GET")
	adminRouter.HandleFunc("/driverequst", controllers.DriveRequest)
	adminRouter.HandleFunc("/blockdriver", controllers.BlockDriver)
	adminRouter.HandleFunc("/blockuser", controllers.BlockUser)
}
