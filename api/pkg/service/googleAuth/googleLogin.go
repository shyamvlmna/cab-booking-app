package googleauth

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/shayamvlmna/cab-booking-app/pkg/models"
	"github.com/shayamvlmna/cab-booking-app/pkg/service/user"
)

var (
	authConfig = &oauth2.Config{
		ClientID:     "662778233746-oeblr3vkk1om82nmjce90lqlac1p7fvq.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/user/googleCallback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	randomState = "randomstate"
)

type AuthContent struct {
	ID            string `json:"id"`
	Firstname     string `json:"given_name"`
	Lastname      string `json:"family_name"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := authConfig.AuthCodeURL(randomState)
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	state := r.URL.Query()["state"][0]
	if state != "randomstate" {

		response := &models.Response{
			ResponseStatus:  "fail",
			ResponseMessage: "states don't match",
			ResponseData:    nil,
		}
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			return
		}
		return
	}
	// if r.FormValue("state") != randomState {
	// 	fmt.Println("not a valid state")
	// 	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	// 	return
	// }
	code := r.URL.Query()["code"][0]
	tok, err := authConfig.Exchange(context.Background(), code)
	if err != nil {
		response := &models.Response{
			ResponseStatus:  "fail",
			ResponseMessage: "code token exange failed",
			ResponseData:    nil,
		}
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			return
		}
		return
	}
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + tok.AccessToken)
	if err != nil {
		response := &models.Response{
			ResponseStatus:  "fail",
			ResponseMessage: "data fetch failed",
			ResponseData:    nil,
		}
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			return
		}
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response := &models.Response{
			ResponseStatus:  "fail",
			ResponseMessage: "json parsing failed",
			ResponseData:    nil,
		}
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			return
		}
		return
	}
	authContent := &AuthContent{}
	if err = json.Unmarshal(content, &authContent); err != nil {
		response := &models.Response{
			ResponseStatus:  "fail",
			ResponseMessage: "unmarshal failed",
			ResponseData:    nil,
		}
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			return
		}
		return
	}

	newUser := &models.User{}

	newUser.Firstname = authContent.Firstname
	newUser.Lastname = authContent.Lastname
	newUser.Email = authContent.Email
	newUser.Picture = authContent.Picture

	user.GoogleAuthUser(newUser)

	// json.NewEncoder(w).Encode(&authContent)

	// http.Redirect(w, r, "/user/signup", http.StatusSeeOther)

	// fmt.Fprintln(w, string(content))
	// {
	// 	"id": "109429758760150763543",
	// 	"email": "shyamvlmna@gmail.com",
	// 	"verified_email": true,
	// 	"name": "Shyamjith P Vilamana",
	// 	"given_name": "Shyamjith",
	// 	"family_name": "P Vilamana",
	// 	"picture": "https://lh3.googleusercontent.com/a-/AOh14Gj4L240leqI64MfmshtoQsqLv_vm0RTPoZ4Z9yCHg=s96-c",
	// 	"locale": "en"
	//   }
	// 	authContent := &AuthContent{}

	// 	json.Unmarshal(content, &authContent)

	// 	log.Println(authContent)
	// 	json.NewEncoder(w).Encode(&authContent)

	// fmt.Println(authContent.Email)
}
