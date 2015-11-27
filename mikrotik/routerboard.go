package mikrotik

import (
	"fmt"

	api "github.com/Netwurx/routeros-api-go"
)

type Resource struct {
	Uptime  string
	Version string
	CpuLoad string
	// CpuLoad uint8
}

type Ethernet struct {
	Name string
}

// chain=input action=accept protocol=tcp src-address-list=adminList in-interface=ether1-gateway dst-port=22
// jump-target=hs-unauth hotspot=from-client,!auth reject-with=tcp-reset
// src-address=194.183.186.0/24 dst-address=194.183.186.0/24
// connection-state=established,related
// connection-state=new connection-nat-state=!dstnat
type FirewallFilter struct {
	ID             uint16
	Name           string
	Chain          string
	Action         string
	Protocol       string
	SrcAddressList string
	DstAddressList string
	SrcAddress     string
	DstAddress     string
	InInterface    string
	DstPort        string
	hotspot        string
	JumpTarget     string
	RejectWith     string
}

// GetResource return mikrotik resource
// free-hdd-space:8380416 architecture-name:smips cpu:MIPS 24Kc V7.4 cpu-frequency:650 total-memory:33554432 bad-blocks:0 version:6.33 (stable) free-memory:7995392 cpu-load:0 total-hdd-space:16777216 platform:MikroTik build-time:Nov/06/2015 12:49:27 cpu-count:1 write-sect-total:10339 board-name:hAP lite uptime:11h45m40s write-sect-since-reboot:865
func GetResource(c *api.Client) (r Resource, err error) {
	res, err := c.Call("/system/resource/getall", nil)
	if err != nil {
		return r, err
	}
	r.Uptime = res.SubPairs[0]["uptime"]
	r.Version = res.SubPairs[0]["version"]
	r.CpuLoad = res.SubPairs[0]["cpu-load"]
	// r. = res.SubPairs[0][""]
	// r. = res.SubPairs[0][""]
	// r. = res.SubPairs[0][""]
	return r, nil
}

func GetInterface(c *api.Client) {
	// var q c.Query
	// q.Pairs = append(q.Pairs, Pair{Key: "type", Value: "ether", Op: "="})

	// res, err := c.Query("/interface/print", q)
	// if err != nil {
	// 	fmt.Printf("GetInterface ERROR: %v\n", err.Error())
	// }
	// if len(res.SubPairs) <= 1 {
	fmt.Println("Did not get multiple SubPairs from interface query")
	// }
	// fmt.Printf("%v\n", res)
}

func GetEthernet(e string, c *api.Client) (eth Ethernet, err error) {
	getEtherAddrs := api.NewPair("interface", e)
	getEtherAddrs.Op = "="
	var q api.Query
	q.Pairs = append(q.Pairs, *getEtherAddrs)
	q.Proplist = []string{"address"}

	res, err := c.Query("/ip/address/print", q)
	if err != nil {
		return eth, err
	}

	for _, v := range res.SubPairs {
		for _, sv := range v {
			fmt.Printf("address: %v\n", sv)
		}
	}
	return eth, nil
}
