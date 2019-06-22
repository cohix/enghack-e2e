package auth

import (
	"fmt"
	"net/http"

	"github.com/cohix/enghack-e2e/server/secret"
	"github.com/cohix/simplcrypto"
)

var token string
var hmacData = "EngHack2019!"

// Verify verifies the authentication hmac
func Verify(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		hmacHeader := r.Header.Get("AUTHORIZATION")

		hmac := simplcrypto.HMACWithSecretAndData(token, hmacData)
		hmacString := simplcrypto.Base64URLEncode(hmac)

		if hmacHeader != hmacString {
			fmt.Printf("unauthorized\n")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func init() {
	token = secret.Generate()
	fmt.Println("export ENGHACKAUTHTOKEN=" + token)
}
