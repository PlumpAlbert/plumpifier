package lib

import (
	"encoding/json"
	"io"
)


func DecodeBody[V any](buffer []byte) (*V, error) {
	var data V
	err := json.Unmarshal(buffer, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
