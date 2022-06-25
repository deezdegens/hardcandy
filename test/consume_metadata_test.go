package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/clog"
)

func TestPingLambda(t *testing.T) {
	URL := "https://pastewin-pastes.s3.us-west-2.amazonaws.com/57ec656e-a14b-4f65-b241-4dfdc02ea5af.json"

	client := clog.New()
	bytes, err := client.Call("GET", URL, nil, map[string]string{})

	fmt.Println(string(bytes), err)
}
