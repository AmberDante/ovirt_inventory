package main

import (
	"encoding/json"
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
	// fmt.Printf("Pass from env: %v\n", ovirtPass)
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

	vms, err := makeVMsStructExportable(VMs)
	if err != nil {
		log.Fatal(err)
	}

	vm_json, err := json.Marshal(vms)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(vm_json), "\n")

}

func makeVMsStructExportable(source []ovirtclient.VM) ([]Vm, error) {
	target := make([]Vm, len(source))
	for i, v := range source {
		target[i].Cpu.Topology.Cores = int(v.CPU().Topo().Cores())
		target[i].Cpu.Topology.Sockets = int(v.CPU().Topo().Sockets())
		target[i].Cpu.Topology.Threads = int(v.CPU().Topo().Threads())
		target[i].Cpu.Mode = CpuMode(fmt.Sprint(v.CPU().Mode()))
		target[i].Comment = v.Comment()
		target[i].Description = v.Description()
		target[i].ID = string(v.ID())
		target[i].Memory = int(v.Memory())
		target[i].MemoryPolicy.Ballooning = v.MemoryPolicy().Ballooning()
		target[i].MemoryPolicy.Max = int(*v.MemoryPolicy().Max())
		target[i].MemoryPolicy.Guaranteed = int(*v.MemoryPolicy().Guaranteed())
		target[i].Name = v.Name()
		target[i].OS.Type = v.OS().Type()
		target[i].SoundcardEnabled = v.SoundcardEnabled()
		target[i].Status = VmStatus(v.Status())
		target[i].Type = VmType(v.VMType())
	}
	return target, nil
}
