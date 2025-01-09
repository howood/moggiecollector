package utils

import (
	"encoding/json"
)

const (
	marshalPrefix = ""
	marshalIndent = "    "
)

func JSONToByte(jsondata interface{}) ([]byte, error) {
	//nolint:wrapcheck
	return json.MarshalIndent(jsondata, marshalPrefix, marshalIndent)
}

func ByteToJSON(jsonbyte []byte) (interface{}, error) {
	var jsondata interface{}
	err := json.Unmarshal(jsonbyte, &jsondata)
	//nolint:wrapcheck
	return jsondata, err
}

func ByteToJSONStruct(jsonbyte []byte, jsonobj interface{}) error {
	//nolint:wrapcheck
	return json.Unmarshal(jsonbyte, &jsonobj)
}
