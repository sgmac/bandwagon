# Bandwagon

A small Go API for BandwagonHOST vps.

## Installation 

The library can be installed the Go way:

``go get -u github.com/sgmac/bandwagon``

## Usage 

Before using the API you need to create a client and set your credentials.

```go
creds := bandwagon.Credentials{
        APIKey: "",
        VeID:   "",
}

c := bandwagon.NewClient(creds)
images, err := c.ListImages()
if err != nil {
        log.Println(err)
}
fmt.Println(images)
```
Below are the API endpoints implemented so far.

### Images

You can list images as in the previous example.

### Server operations

These are the basic operations that you can perform on a server.

**Start**
```go
c := bandwagon.NewClient(creds)
resp, err := c.Start()
```

**Stop**
```go
c := bandwagon.NewClient(creds)
resp, err := c.Stop()
```

**Reboot**
```go
c := bandwagon.NewClient(creds)
resp, err := c.Reboot()
```

**Install OS**
```go
c := bandwagon.NewClient(creds)
resp, err := c.Install("ubuntu-13.10-x86_64")
```

**Set hostname**
```go
c := bandwagon.NewClient(creds)
resp, err := c.Hostname("ragnarok")
```

**Get info**

This is going to get information for the VPS running. Returns a struct of type Info:

```go
type InfoVPS struct {
        VMType                string            `json:"vm_type"`
        Hostname              string            `json:"hostname"`
        NodeIP                string            `json:"node_ip"`
        NodeAlias             string            `json:"node_alias"`
        NodeLocation          string            `json:"node_location"`
        LocationIPv6Ready     bool              `json:"location_ipv6_ready"`
        Plan                  string            `json:"plan"`
        PlanMonthlyData       int64             `json:"plan_monthly_data"`
        MonthlyDataMultiplier int64             `json:"plan_monthly_data"`
        PlanDisk              int64             `json:"plan_disk"`
        PlanRAM               int32             `json:"plan_ram"`
        PlanSwap              int32             `json:"plan_swap"`
        PlanMaxIPv6           int32             `json:"plan_max_i_pv_6"`
        OS                    string            `json:"os"`
        Email                 string            `json:"email"`
        DataCounter           int32             `json:"data_counter"`
        DataNextReset         int32             `json:"data_next_reset"`
        IPAddresses           []string          `json:"ip_addresses"`
        RDNSApiAvailable      bool              `json:"rdns_api_available"`
        PTR                   map[string]string `json:"ptr"`
        Suspended             bool              `json:"suspended"`
        Error                 int32             `json:"error"`
}
```

Similar to previous API calls:

```go
c := bandwagon.NewClient(creds)
resp, err := c.Info()
```

## License

MIT License

Copyright (c) [2017] [Sergio Galvan]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
