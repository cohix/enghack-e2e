package action

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"

	"github.com/cohix/simplcrypto"
)

//GetOrCreateKey gets a key from disk if it exists, or creates a new key if not
func GetOrCreateKey() (*simplcrypto.SymKey, error) {
	keypath := keypath()
	if keypath == "" {
		return nil, errors.New("failed to get keypath")
	}

	var key *simplcrypto.SymKey

	key, err := keyFromFile(keypath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get keyFromFile")
	}

	if key != nil {
		return key, nil
	}

	key, err = simplcrypto.GenerateSymKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to GenerateSymKey")
	}

	if err := writeKeyToFile(keypath, key); err != nil {
		return nil, errors.Wrap(err, "failed to writeKeyToFile")
	}

	return key, nil
}

func keyFromFile(filepath string) (*simplcrypto.SymKey, error) {
	keyFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var key simplcrypto.SymKey
	if err := json.Unmarshal(keyFile, &key); err != nil {
		return nil, err
	}

	return &key, nil
}

func writeKeyToFile(filepath string, key *simplcrypto.SymKey) error {
	symKeyJSON, err := json.MarshalIndent(key, "", "\t")
	if err != nil {
		return errors.Wrap(err, "failed to Marshal key")
	}

	if err := ioutil.WriteFile(filepath, symKeyJSON, 0600); err != nil {
		return errors.Wrap(err, "failed to WriteFile")
	}

	return nil
}

func keypath() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	return filepath.Join(homedir, ".enghack-key.json")
}
