// Package JSONApiFormatter provides ...
package JSONApiFormatter

import (
	"encoding/json"
)

func Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}
