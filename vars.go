package main

import "time"

type vmDisks struct {
	vmID            string
	diskAttachments []DiskAttachment
}

type vmStats struct {
	Comment string `json:"comment,omitempty"` // Free text containing comments about this object.
	Cpu     int    `json:"cpu,omitempty"`     // Numder of CPU for the virtual machine.
	// The virtual machine creation date.
	// When requesting the JSON representation the engine uses a different, format:
	// an integer containing the number of seconds since Jan 1st 1970, also know as
	// epoch time [https://en.wikipedia.org/wiki/Unix_time].
	CreationTime time.Duration `json:"creation_time,omitempty"`
	Description  string        `json:"description,omitempty"`   // A human-readable description in plain text.
	FQDN         string        `json:"fqdn,omitempty"`          // Fully qualified domain name of the virtual machine.
	ID           string        `json:"id,omitempty"`            // A unique identifier.
	Memory       int           `json:"memory,omitempty"`        // The virtual machine’s memory, in bytes.
	Name         string        `json:"name,omitempty"`          // A human-readable name in plain text.
	OS           string        `json:"os,omitempty"`            // Operating system type installed on the virtual machine. From the field Type from struct OperatingSystem
	RunOnce      bool          `json:"run_once,omitempty"`      // If true, the virtual machine has been started using the run once command, meaning it’s configuration might differ from the stored one for the purpose of this single run.
	SerialNumber string        `json:"serial_number,omitempty"` // Virtual machine’s serial number in a cluster. The field Value of struct SerialNumber
	StartTime    time.Duration `json:"start_time,omitempty"`    // The date in which the virtual machine was started.
	Status       VmStatus      `json:"status,omitempty"`        // The current status of the virtual machine.
	StatusDetail string        `json:"status_detail,omitempty"` // Human readable detail of current status.
	StopReason   string        `json:"stop_reason,omitempty"`   // The reason the virtual machine was stopped.
	StopTime     time.Duration `json:"stop_time,omitempty"`     // The date in which the virtual machine was stopped.

	HddDiskSize  int `json:"hdd_disk_size,omitempty"`  // The size of all disks attached to vm. Sum of initial_size from Disk struct, in bytes. Group by StorageType (HDD or SSD).
	SsdDiskSize  int `json:"ssd_disk_size,omitempty"`  // The size of all disks attached to vm. Sum of initial_size from Disk struct, in bytes. Group by StorageType (HDD or SSD).
	VmDisksCount int `json:"vm_disks_count,omitempty"` // count of attached virtual disks
}
