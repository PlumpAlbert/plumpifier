package lib

import (
	"bytes"
	"encoding/json"
	"text/template"

)

func DecodeBody[V any](buffer []byte) (*V, error) {
	var data V
	err := json.Unmarshal(buffer, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func PrepareTemplate(tpl string, data any) (string, error) {
	msg, err := template.New("msg").Parse(tpl)
	if err != nil {
		return "", err
	}

	var tmp bytes.Buffer
	msg.Execute(&tmp, data)

	return tmp.String(), nil
}
