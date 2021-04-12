package authenticationhandler

import (
	"encoding/json"
	"fmt"
	authentication "hcn/myhelpers/authenticationHelper"
	token "hcn/myhelpers/tokenHelper"
	"hcn/mymodels"
	"io/ioutil"
	"net/http"
)

// Login endpoint receives username and password and validate them.
// Returns token if ok, othwerwise an error
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var teacher mymodels.Teacher
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	json.Unmarshal(reqBody, &teacher)
	switch {
	case (teacher.Email == nil) || (len(*teacher.Email) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Email is empty or not valid")
		fmt.Fprintf(w, "Email is empty or not valid")
		return
	case (teacher.Password == nil) || (len(*teacher.Password) == 0):
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Password is empty or not valid")
		fmt.Fprintf(w, "Password is empty or not valid")
		return
	default:
		token, err := authentication.UserAuthentication(teacher)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(token)
		}
		return
	}
}

// IsValid endpoint receives a token and checks if is valid
func IsValid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var teacherToken mymodels.Token
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "(USER) %v", err.Error())
		return
	}
	json.Unmarshal(reqBody, &teacherToken)
	_, err = token.VerifyAuthenticity(*teacherToken.Token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(err.Error())
	}
	return
}
