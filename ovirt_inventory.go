package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

const ovirtUrl = "ovirttest01.pkcht.ru"
const apiURI = "/ovirt-engine/api"

func main() {
	var user string = "admin@internal"
	var scope = []string{"ovirt-app-api"}
	var ovirtAPIurl string = ovirtUrl + apiURI
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     user,
		ClientSecret: os.Getenv("OVIRT_PASS"),
		Scopes:       scope,
		Endpoint: oauth2.Endpoint{
			AuthURL:  ovirtUrl,
			TokenURL: ovirtUrl,
		},
	}

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the custom HTTP client when requesting a token.
	tlsClient := &http.Client{Transport: tlsTranstort()}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, tlsClient)
	token, err := conf.PasswordCredentialsToken(ctx, conf.ClientID, conf.ClientSecret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("access_token: %v\nTokenType: %v\nExpiry: %v", token.AccessToken, token.TokenType, token.Expiry)
	client := conf.Client(ctx, token)

	vms, err := vmsList(client, ovirtAPIurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(vms)

	vmDiskList, err := vmsDisks(client, ovirtAPIurl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(vmDiskList)

	disksForVms, err := diskAttachments(client, ovirtAPIurl, vms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(disksForVms)

	vmsStats := composeStats(vms, vmDiskList, disksForVms)
	fmt.Print(vmsStats)
}

func getRequest(client *http.Client, url string) ([]byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}

func vmsList(client *http.Client, url string) ([]Vm, error) {

	url = url + "/vms"

	body, err := getRequest(client, url)
	if err != nil {
		return nil, err
	}
	var vms []Vm
	if err := json.Unmarshal(body, &vms); err != nil {
		return nil, err
	}
	return vms, nil
}

func vmsDisks(client *http.Client, url string) ([]Disk, error) {
	url = url + "/vmdisks"
	body, err := getRequest(client, url)
	if err != nil {
		return nil, err
	}
	var vmdisks []Disk
	if err := json.Unmarshal(body, &vmdisks); err != nil {
		return nil, err
	}
	return vmdisks, nil
}

func diskAttachments(client *http.Client, url string, vms []Vm) ([]vmDisks, error) {

	var vmsDisksAttachements []vmDisks
	for _, vm := range vms {
		url = url + "/vms/" + vm.ID + "/diskattachments"
		body, err := getRequest(client, url)
		if err != nil {
			return nil, err
		}
		var attachments []DiskAttachment
		if err := json.Unmarshal(body, &attachments); err != nil {
			return nil, err
		}

		vmsDisksAttachements = append(vmsDisksAttachements, vmDisks{vmID: vm.ID, diskAttachments: attachments})
	}
	return vmsDisksAttachements, nil
}

func tlsTranstort() *http.Transport {
	return &http.Transport{
		TLSClientConfig: tlsConfig(),
	}
}

func tlsConfig() *tls.Config {

	crt, err := os.ReadFile("./oVirt_CA/ovirt_ca.pem")
	if err != nil {
		log.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         ovirtUrl,
	}
}

func composeStats(vms []Vm, diskList []Disk, disksForVms []vmDisks) []vmStats {
	var vmsStats = make([]vmStats, 0, len(vms))
	// Create Map of Disk
	disksMap := diskMap(diskList)
	// Create Map for attached disks to VMs
	vmsDisksMap := vmDisksMap(disksForVms)

	for i, v := range vms {
		vmsStats[i].Comment = v.Comment
		// TODO: make a request to oVirt API and check out how Vm.Cpu.Cores looks like.
		// vmsStats[i].Cpu = v.Cpu.Cores
		vmsStats[i].CreationTime = v.CreationTime
		vmsStats[i].Description = v.Description
		vmsStats[i].FQDN = v.FQDN
		vmsStats[i].ID = v.ID
		vmsStats[i].Memory = v.Memory
		vmsStats[i].Name = v.Name
		vmsStats[i].OS = v.OS.Type
		vmsStats[i].RunOnce = v.RunOnce
		vmsStats[i].SerialNumber = v.StatusDetail
		vmsStats[i].StartTime = v.StartTime
		vmsStats[i].Status = v.Status
		vmsStats[i].StatusDetail = v.StatusDetail
		vmsStats[i].StopReason = v.StopReason
		vmsStats[i].StopTime = v.StopTime

		vmsStats[i].HddDiskSize = diskSizeCalculation(vmsDisksMap[v.ID], disksMap, false)
		vmsStats[i].SsdDiskSize = diskSizeCalculation(vmsDisksMap[v.ID], disksMap, true)
		vmsStats[i].VmDisksCount = len(vmsDisksMap[v.ID])

	}
	return vmsStats
}

// Prepare Map for easiest search attached disks to VM
func vmDisksMap(disksForVms []vmDisks) map[string][]DiskAttachment {
	var vmsDisksMap = make(map[string][]DiskAttachment)
	for _, disksForVm := range disksForVms {
		vmsDisksMap[disksForVm.vmID] = disksForVm.diskAttachments
	}
	return vmsDisksMap
}

// Prepare Map for []Disk
func diskMap(diskList []Disk) map[string]Disk {
	var disksMap = make(map[string]Disk, 0)
	for _, disk := range diskList {
		disksMap[disk.ID] = disk
	}
	return disksMap
}

// func diskSizeCalculation - calculate disks size grouped by StorageType HDD or SSD
//
//   - attachedDisks - Slise of attacet disks to VM
//   - disks - Map of virtual disks
//   - ssd
//
// if ssdStorage == ssd - sum will be calculated.
func diskSizeCalculation(attachedDisks []DiskAttachment, disks map[string]Disk, ssd bool) int {
	var diskSize int
	for _, v := range attachedDisks {
		// calculate discs size
		if ssdStorage(disks[v.ID].LunStorage.Description) == ssd {
			diskSize += disks[v.ID].InitialSize
		}

	}
	return diskSize
}

// func ssdStorage - helps to determine is virtual disk on SSD storage or not
func ssdStorage(description string) bool {

	return strings.Contains(strings.ToLower(description), "ssd")
}
