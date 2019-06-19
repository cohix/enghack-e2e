package auth

import (
	"fmt"
	"net/http"

	"github.com/cohix/simplcrypto"
	"github.com/enghack-e2e/server/secret"
)

var enableAuth = true

// Key defines the key used for auth
var Key string

// Verify verifies the authentication hmac
func Verify(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if enableAuth {
			hmacHeader := r.Header.Get("AUTHORIZATION")
			hmac := simplcrypto.Base64URLEncode(simplcrypto.HMACWithSecretAndData(Key, "EngHack2019!"))

			if hmacHeader != hmac {
				fmt.Printf("unauthorized\n")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}

func init() {
	Key = secret.Generate()
	fmt.Println("export ENGHACKAUTHKEY=" + Key)
}
