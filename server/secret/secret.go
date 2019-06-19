package secret

import (
	cryptorand "crypto/rand"

	"github.com/cohix/simplcrypto"
)

// Generate generates a random 32 character secret
func Generate() string {
	b := make([]byte, 32)
	if _, err := cryptorand.Read(b); err != nil {
		panic(err)
	}

	return simplcrypto.Base64URLEncode(b)
	//return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}
