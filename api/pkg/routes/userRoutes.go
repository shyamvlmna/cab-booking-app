package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shayamvlmna/cab-booking-app/pkg/controllers"
	"github.com/shayamvlmna/cab-booking-app/pkg/middleware"
	"github.com/shayamvlmna/cab-booking-app/pkg/models"
	googleauth "github.com/shayamvlmna/cab-booking-app/pkg/service/googleAuth"
)

func UserRoutes(r *mux.Router) {

	r.HandleFunc("/", controllers.Index)

	userRouter := r.PathPrefix("/user").Subrouter()

	//check wheather phonenumber already registerd or is a new entry
	userRouter.HandleFunc("/auth", controllers.UserAuth).Methods("POST")

	userRouter.HandleFunc("/signup", controllers.UserSignUp).Methods("POST")
	userRouter.HandleFunc("/login", controllers.UserLogin).Methods("POST")
	userRouter.HandleFunc("/googlelogin", googleauth.GoogleLogin)
	userRouter.HandleFunc("/googleCallback", googleauth.GoogleCallback)
	userRouter.HandleFunc("/logout", controllers.UserLogout)

	//render enter otp page
	userRouter.HandleFunc("/enterotp", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := &models.Response{
			ResponseStatus:  "success",
			ResponseMessage: "new user",
			ResponseData:    nil,
		}
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			return
		}

	}).Methods("GET")

	//validate submited otp
	userRouter.Handle("/otp", middleware.ValidateOtp(controllers.UserSignupPage)).Methods("POST")

	//render login page to enter password
	userRouter.HandleFunc("/loginpage", controllers.UserLoginPage).Methods("GET")

	//render the homepage only if authorized with JWT
	userRouter.Handle("/userhome", middleware.IsAuthorized(controllers.UserHome)).Methods("GET")

	userRouter.HandleFunc("/update/{id}", controllers.EditUserProfile).Methods("GET")
	userRouter.HandleFunc("/update/{id}", controllers.UpdateUserProfile).Methods("POST")

	//book new trip
	userRouter.Handle("/booktrip", middleware.IsAuthorized(controllers.BookTrip)).Methods("POST")

	userRouter.Handle("/triphistory", middleware.IsAuthorized(controllers.UserTripHistory)).Methods("GET")

	// userRouter.HandleFunc("/test", controllers.Test)

	userRouter.Handle("/confirmtrip", middleware.IsAuthorized(controllers.ConfirmTrip)).Methods("POST")

}
