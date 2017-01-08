package cmn4j_test

import (
	"testing"

	"github.com/kc1116/cmool-neo4j/cmn4j"
)

var api cmn4j.API

func TestCreateEventNode(t *testing.T) {
	uri := "bs"
	_, err := api.Connect(uri)

	f, ok := err.(error)
	if !ok {
		t.Error("Expected an error got", f)
	}

}
