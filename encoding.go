package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
)

// JSONFileToStruct encodes json file to struct
func JSONFileToStruct(jsonFile string, out interface{}) error {
	if reflect.ValueOf(out).Kind() != reflect.Ptr {
		return errors.New("destination type not pointer")
	}

	fmt.Println(out)

	raw, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, &out)
}

// ByteArrayToString converts byte array into string
func ByteArrayToString(bs []uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}
