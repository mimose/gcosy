package lib

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

func Encode(data interface{}) ([]byte, error) {
	var jsonBuffer bytes.Buffer
	jsonEncoder := json.NewEncoder(&jsonBuffer)
	err := jsonEncoder.Encode(&data)
	if err != nil {
		return nil, err
	}

	var gobBuffer bytes.Buffer
	gobEncoder := gob.NewEncoder(&gobBuffer)
	err = gobEncoder.Encode(jsonBuffer.Bytes())
	if err != nil {
		return nil, err
	}
	return gobBuffer.Bytes(), err
}
