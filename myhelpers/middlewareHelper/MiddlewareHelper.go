package middlewarehelper

import (
	"encoding/json"
	"fmt"
	token "hcn/myhelpers/tokenHelper"

	"bytes"
	"io/ioutil"
	"net/http"
)

// Middleware for checking user authorization
func Middleware(handler http.HandlerFunc, endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println("IP -> ", originIP, "Endpoint ->", endpoint, "        isAuth:", isAuth, err)
		if !isAuth {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Middleware authentication error: " + err.Error())
			return
		}
		handler.ServeHTTP(w, r)
		return
	}
}
