package test

import (
	"fmt"
	"testing"

	"github.com/deezdegens/clog"
)

func TestPingLambda(t *testing.T) {
	URL := "https://pastewin-pastes.s3.us-west-2.amazonaws.com/b3c6ec23-58c0-4875-a709-3a51c3257be5.json"

	client := clog.New()
	bytes, err := client.Call("GET", URL, nil, map[string]string{})

	fmt.Println(string(bytes), err)
}
