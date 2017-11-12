package bandwagon

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestListImages(t *testing.T) {

	var imagesResponse = `
	{
	"error": 0,
	"installed": "centos",
	"templates": ["debian-x64", "centos-7" ]
	}
	`
	setup()
	mux.HandleFunc("/v1/getAvailableOS", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, imagesResponse)
	})

	images, err := client.ListImages()
	if err != nil {
		t.Errorf("client.ListImages got %v\n err :%v\n", images, err)
	}
	expected := &Images{
		Templates: []string{"debian-x64", "centos-7"},
		Installed: "centos",
		Error:     0,
	}

	if !reflect.DeepEqual(images, expected) {
		t.Errorf("client.ListImages got %v expected %v\n", images, expected)
	}

}

func TestListImagesDecode(t *testing.T) {

	var forceError = `[]`

	setup()
	mux.HandleFunc("/v1/getAvailableOS", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, forceError)
	})

	_, err := client.ListImages()
	if err != nil {
		t.Logf("client.ListImages could not Unmarshal, got :%v\n", forceError)
	}
}
