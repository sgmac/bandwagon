package bandwagon

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
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
	fmt.Println(mux)

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
		t.Errorf("client.Kill got %v\n err :%v\n", resp, err)
	}

	log.Println("With a success kill we should get")
	{
		log.Printf("\t%c %d error response, expected 0\n", success, resp.Error)
	}
}

func TestServerInfo(t *testing.T) {

	var info = `
	{
	  "vm_type": "ovz",
	  "hostname": "lab",
	  "node_ip": "35.12.220.103",
	  "node_alias": "v903",
	  "node_location": "EU, Netherlands",
	  "location_ipv6_ready": true,
	  "plan": "wagon5promo",
	  "plan_monthly_data": 0,
	  "monthly_data_multiplier": 0,
	  "plan_disk": 1,
	  "plan_ram": 1,
	  "plan_swap": 1,
	  "plan_max_ipv6s": 0,
	  "os": "centos-7-x86_64-minimal",
	  "email": "some@mail.com",
	  "data_counter": 1,
	  "data_next_reset": 1,
	  "ip_addresses": [
	    "162.113.12.78"
	  ],
	  "rdns_api_available": true,
	  "suspended": false,
	  "error": 0
	}
	`

	setup()
	mux.HandleFunc("/v1/getServiceInfo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, info)
	})

	response, err := client.Info()
	if err != nil {
		t.Errorf("client.Info got %v\n err :%v\n", response, err)
	}

	expected := &InfoVPS{
		VMType:                "ovz",
		Hostname:              "lab",
		NodeIP:                "35.12.220.103",
		NodeAlias:             "v903",
		NodeLocation:          "EU, Netherlands",
		LocationIPv6Ready:     true,
		Plan:                  "wagon5promo",
		PlanMonthlyData:       0,
		MonthlyDataMultiplier: 0,
		PlanDisk:              1,
		PlanRAM:               1,
		PlanSwap:              1,
		PlanMaxIPv6:           0,
		OS:                    "centos-7-x86_64-minimal",
		Email:                 "some@mail.com",
		DataCounter:           1,
		DataNextReset:         1,
		IPAddresses:           []string{"162.113.12.78"},
		RDNSApiAvailable:      true,
		Suspended:             false,
		Error:                 0,
	}

	log.Println("When getting info from a VPS")
	{
		if !reflect.DeepEqual(response, expected) {
			t.Errorf("\t %c tclient.Info got %v\n expected %v\n", failure, response, expected)

		}

		log.Printf("\t %c response data should be valid \n", success)
	}

}
