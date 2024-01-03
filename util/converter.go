package util

import (
	"encoding/json"
)

// Converting data to byte in string type
func ToByteString(data any) (string, error) {
	// Encoding data
	var b, err = json.Marshal(&data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
