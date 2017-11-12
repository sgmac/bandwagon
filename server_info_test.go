package bandwagon

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

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
