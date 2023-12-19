package lib

import (
	"encoding/json"
	"io"
)

func DecodeBody[V any](body io.ReadCloser) (*V, error) {
	decoder := json.NewDecoder(body)

	var data V
	err := decoder.Decode(&data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
