package helpers

import (
	"hcn/config"
	"net/http"
	"strings"
)

// VerifyRequest verifies if the client IP
// is the same IP of the server.
func VerifyRequest(r *http.Request) bool {
	// You can add some IP adressess if you need it.
	addresses := [...]string{config.ServerIP}
	for _, address := range addresses {
		if address == strings.Split(r.RemoteAddr, ":")[0] {
			return true
		}
	}
	return false
}
