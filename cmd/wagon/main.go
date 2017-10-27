package main

import (
	"fmt"
	"log"

	"github.com/sgmac/bandwagon"
)

func main() {

	creds := bandwagon.Credentials{
		APIKey: "",
		VeID:   "",
	}

	c := bandwagon.NewClient(creds)
	data, err := c.ListImages()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
