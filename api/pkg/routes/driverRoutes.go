package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shayamvlmna/cab-booking-app/pkg/controllers"
	"github.com/shayamvlmna/cab-booking-app/pkg/middleware"
)

func DriverRoutes(r *mux.Router) {
	driverRouter := r.PathPrefix("/driver").Subrouter()

	driverRouter.HandleFunc("/auth", controllers.DriverAuth).Methods("POST")
	driverRouter.HandleFunc("/otp", controllers.ValidateDriverOtp).Methods("POST")
	driverRouter.HandleFunc("/signup", controllers.DriverSignUp).Methods("POST")
	driverRouter.HandleFunc("/login", controllers.DriverLogin).Methods("POST")
	driverRouter.HandleFunc("/logout", controllers.DriverLogout).Methods("GET")

	driverRouter.Handle("/driverhome", middleware.IsAuthorized(controllers.DriverHome)).Methods("GET")
	driverRouter.Handle("/regtodrive", middleware.IsAuthorized(controllers.RegisterDriver)).Methods("POST")
	driverRouter.Handle("/addcab", middleware.IsAuthorized(controllers.AddCab)).Methods("POST")

	//render enter otp page
	driverRouter.HandleFunc("/enterotp", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("submit the otp"))
	}).Methods("GET")

	//render signup page
	driverRouter.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("driver signup page\n\nfirstname\nlastname\nemail\ncity\nlicence number\npassword"))
	}).Methods("GET")

	//enter login page to enter password
	driverRouter.HandleFunc("/loginpage", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login page\nOnly submit the password"))
	}).Methods("GET")

}
