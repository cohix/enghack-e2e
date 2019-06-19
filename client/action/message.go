package action

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/cohix/simplcrypto"
	"github.com/pkg/errors"
)

var client = http.DefaultClient

// SetMessage sets the message on the server
func SetMessage(message *simplcrypto.Message) error {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "failed to Marshal message")
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/message", bytes.NewBuffer(messageJSON))
	if err != nil {
		return errors.Wrap(err, "failed to NewRequest")
	}

	hmac, err := getHMACdToken()
	if err != nil {
		return errors.Wrap(err, "failed to getHMACdToken")
	}

	req.Header.Set("AUTHORIZATION", hmac)

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to Do request")
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("set message request failed with status code %s", resp.Status)
	}

	return nil
}

// GetMessage fetches the message from the server
func GetMessage() (*simplcrypto.Message, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/message", nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to NewRequest")
	}

	hmac, err := getHMACdToken()
	if err != nil {
		return nil, errors.Wrap(err, "failed to getHMACdToken")
	}

	req.Header.Set("AUTHORIZATION", hmac)

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to Do request")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get message request failed with status code %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, errors.Wrap(err, "failed to ReadAll response body")
	}

	var message simplcrypto.Message
	if err := json.Unmarshal(body, &message); err != nil {
		return nil, errors.Wrap(err, "failed to Unmarshal message")
	}

	return &message, nil
}

const hmacData = "EngHack2019!"

func getHMACdToken() (string, error) {
	token := os.Getenv("ENGHACKAUTHKEY")
	if token == "" {
		return "", errors.New("missing ENGHACKAUTHKEY environment variable")
	}

	hmac := simplcrypto.HMACWithSecretAndData(token, hmacData)

	hmacString := simplcrypto.Base64URLEncode(hmac)

	return hmacString, nil
}
