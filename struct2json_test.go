// Package JSONApiFormatter provides ...
package JSONApiFormatter

import (
	"bytes"
	"testing"
)

type MyObject struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TestEncodeBasicObject(t *testing.T) {
	want := []byte(`{"myobject":{"id":1,"title":"My title"}}`)
	obj := MyObject{ID: 1, Title: "My title"}
	got, err := Encode(obj)
	if err != nil {
		t.Error("Error while encoding object")
	}
	if !bytes.Equal(got, want) {
		t.Logf("got : %v", string(got))
		t.Logf("want: %v", string(want))
		t.Error("Encoding does not match decoding")
	}
}
