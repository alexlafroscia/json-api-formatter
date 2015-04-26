// Package JSONApiFormatter provides ...
package JSONApiFormatter

import (
	"bytes"
	"encoding/json"
	//"github.com/acsellers/inflections"
	"reflect"
	"strings"
)

func Encode(obj interface{}) ([]byte, error) {
	kind := reflect.TypeOf(obj).Kind()

	switch kind {
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		encodeArray(obj)
	case reflect.Struct:
		encodeSingle(obj)
	default:
		panic("Unrecognized type")
	}

	return encodeSingle(obj)
}

func encodeSingle(obj interface{}) ([]byte, error) {
	objType := reflect.TypeOf(obj).Name()
	objType = strings.ToLower(objType)
	return encodeWithName(obj, objType)
}

func encodeArray(obj interface{}) ([]byte, error) {
	objType := reflect.TypeOf(obj).Name()
	objType = strings.ToLower(objType)
	//objType = inflections.Pluralize(objType)
	return encodeWithName(obj, objType)
}

func encodeWithName(obj interface{}, name string) ([]byte, error) {
	marshalized, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	toReturn := [][]byte{
		[]byte(`{"`),
		[]byte(name),
		[]byte(`":`),
		marshalized,
		[]byte(`}`),
	}
	return bytes.Join(toReturn, nil), nil
}
