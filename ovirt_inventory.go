package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	ovirtclientlog "github.com/ovirt/go-ovirt-client-log/v3"
	ovirtclient "github.com/ovirt/go-ovirt-client/v3"
)

const ovirtUrl = "ovirttest01.pkcht.ru"
const apiUrl = "/ovirt-engine/api"

// const tokenURL = "/ovirt-engine/sso/oauth/token"

type OVirttInfo struct {
	ovirtclient.ClientWithLegacySupport
}

func main() {
	var user string = "admin@internal"
	var ovirtAPIurl string = "https://" + ovirtUrl + apiUrl
	ovirtPass := os.Getenv("OVIRT_PASS")
	fmt.Printf("Pass from env: %v\n", ovirtPass)
	// Create a logger that logs to the standard Go log here:
	// logger := ovirtclientlog.NewGoLogger()

	// Create an ovirtclient.TLSProvider implementation. This allows for simple
	// TLS configuration.
	tls := ovirtclient.TLS()
	tls.CACertsFromDir(
		"./oVirt_CA",
		regexp.MustCompile(`\.pem`),
	)

	client, err := ovirtclient.New(
		// URL to your oVirt engine API here:
		ovirtAPIurl,
		// Username here:
		user,
		// Password here:
		ovirtPass,
		// Pass the TLS provider here:
		tls,
		// Pass the logger here:
		ovirtclientlog.NewNOOPLogger(),
		// Pass in extra settings here. Must implement the ovirtclient.ExtraSettings interface.
		nil,
	)
	if err != nil {
		log.Fatal(err)

	}
	VMs, err := client.ListVMs(ovirtclient.AutoRetry())
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range VMs {
		fmt.Printf("%+v", v)
	}

}
