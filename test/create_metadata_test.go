package test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/deezdegens/clog"
	"github.com/deezdegens/hardcandy"
)

func TestCreateMetadata(t *testing.T) {
	URL := "https://jzd5pgu3q3axdvxbw6kupn3bg40wbnkz.lambda-url.us-west-2.on.aws/"

	metadata := hardcandy.Metadata{
		Image:       "https://images.unsplash.com/photo-1529778873920-4da4926a72c2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Y3V0ZSUyMGNhdHxlbnwwfHwwfHw%3D&w=1000&q=80",
		Description: "What a cute kitty",
		Properties: hardcandy.MetadataProperties{
			Category: "image",
			Files: []hardcandy.MetadataFile{
				{
					URI:  "https://images.unsplash.com/photo-1529778873920-4da4926a72c2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Y3V0ZSUyMGNhdHxlbnwwfHwwfHw%3D&w=1000&q=80",
					Type: "image/jpg",
				},
			},
		},
		Attributes: []hardcandy.MetadataAttribute{
			{
				TraitType: "Silly Goose",
				Value:     "You KNOW IT",
			},
			{
				TraitType: "Piggly Wiggly",
				Value:     "DUH!!!",
			},
		},
	}

	body, _ := json.Marshal(metadata)

	client := clog.New()
	bytes, err := client.Call("POST", URL, body, map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})

	fmt.Println(string(bytes), err)

}
