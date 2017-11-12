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
		log.Printf("\t%c %d error response, expected 0\n", success, resp.Error)
	}

}

func TestFailedReboot(t *testing.T) {

	var failedReboot = `
	{
	"error": 1 
	}
	`
	setup()
	mux.HandleFunc("/v1/restart", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, failedReboot)
	})

	resp, err := client.Reboot()
	if err != nil {
		t.Errorf("client.Reboot got %v\n err :%v\n", resp, err)
	}

	log.Println("With a failed reboot we should get")
	{
		log.Printf("\t%c %d error response, expected 1\n", failure, resp.Error)
	}

}

func TestSuccessStart(t *testing.T) {

	var successStart = `
	{
	"error": 0
	}
	`
	setup()
	mux.HandleFunc("/v1/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, successStart)
	})

	resp, err := client.Start()
	if err != nil {
		t.Errorf("client.Start got %v\n err :%v\n", resp, err)
	}

	log.Println("With a success start we should get")
	{
		log.Printf("\t%c %d error response, expected 0\n", success, resp.Error)
	}
}

func TestSuccessStop(t *testing.T) {

	var successStop = `
	{
	"error": 0
	}
	`
	setup()
	mux.HandleFunc("/v1/stop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, successStop)
	})

	resp, err := client.Stop()
	if err != nil {
		t.Errorf("client.Stop got %v\n err :%v\n", resp, err)
	}

	log.Println("With a success stop we should get")
	{
		log.Printf("\t%c %d error response, expected 0\n", success, resp.Error)
	}
}

func TestSuccessKill(t *testing.T) {

	var successKill = `
	{
	"error": 0
	}
	`
	setup()
	mux.HandleFunc("/v1/kill", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, successKill)
	})

	resp, err := client.Kill()
	if err != nil {
		t.Errorf("client.Stop got %v\n err :%v\n", resp, err)
	}

	log.Println("With a success kill we should get")
	{
		log.Printf("\t%c %d error response, expected 0\n", success, resp.Error)
	}
}
