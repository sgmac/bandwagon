package bandwagon

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

const success = '\u2713'
const failure = '\u2717'

func TestSuccessReboot(t *testing.T) {

	var successReboot = `
	{
	"error": 0
	}
	`
	setup()
	mux.HandleFunc("/v1/restart", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, successReboot)
	})

	resp, err := client.Reboot()
	if err != nil {
		t.Errorf("client.Reboot got %v\n err :%v\n", resp, err)
	}

	log.Println("With a success reboot we should get")
	{
		log.Printf("\t%c got %d error response, expected 0\n", success, resp.Error)
	}

}

func TestFailedReboot(t *testing.T) {

	var successReboot = `
	{
	"error": 1 
	}
	`
	setup()
	mux.HandleFunc("/v1/restart", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, successReboot)
	})

	resp, err := client.Reboot()
	if err != nil {
		t.Errorf("client.Reboot got %v\n err :%v\n", resp, err)
	}

	log.Println("With a success reboot we should get")
	{
		log.Printf("\t%c got %d error response, expected 1\n", failure, resp.Error)
	}

}
