package middlewareHelper

import (
	"encoding/json"
	"fmt"
	token "hcn/myhelpers/tokenHelper"

	"bytes"
	"io/ioutil"
	"net/http"
)

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	fmt.Println("Erdaaa")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Middleware for checking user authorization
func Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Ajaaa 1")
		setupResponse(&w, r)
		fmt.Println("Ajaaa 2", (*r).Method)
		if (*r).Method == "OPTIONS" {
			fmt.Println("Ajaaa 3")
			return
		}
		fmt.Println("Ajaaa 4")

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "(USER) %v", err.Error())
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

		originIP := r.Header.Get("X-FORWARDED-FOR")
		if originIP != "" {
			originIP = ""
		} else {
			originIP = r.RemoteAddr
		}
		fmt.Println("A", r.Header.Get("Token"))
		isAuth, err := token.IsValid(r.Header.Get("Token"))
		fmt.Println("IP ->", originIP, "  Endpoint ->", r.RequestURI)
		if !isAuth {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Middleware authentication error: " + err.Error())
			return
		}
		handler.ServeHTTP(w, r)
		return
	}
}
