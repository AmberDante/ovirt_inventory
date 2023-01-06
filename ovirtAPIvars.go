package main

import "time"

type Vm struct {
	// Reference to virtual machine’s BIOS configuration.
	Bios Bios `json:"bios,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// Console configured for this virtual machine.
	Console Console `json:"console,omitempty"`
	// The configuration of the virtual machine CPU.
	Cpu       Cpu `json:"cpu,omitempty"`
	CpuShares int `json:"cpu_shares,omitempty"`
	// The virtual machine creation date.
	// When requesting the JSON representation the engine uses a different, format:
	// an integer containing the number of seconds since Jan 1st 1970, also know as epoch time [https://en.wikipedia.org/wiki/Unix_time].
	CreationTime               time.Duration    `json:"creation_time"`
	CustomCompatibilityVersion Version          `json:"custom_compatibility_version,omitempty"`
	CustomCpuModel             string           `json:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine      string           `json:"custom_emulated_machine,omitempty"`
	CustomProperty             []CustomProperty `json:"custom_property,omitempty"`
	// If true, the virtual machine cannot be deleted.
	DeleteProtected bool `json:"delete_protected,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// The virtual machine display configuration.
	Display Display `json:"display,omitempty"`
	// Domain configured for this virtual machine.
	Domain Domain `json:"domain,omitempty"`
	// Fully qualified domain name of the virtual machine.
	FQDN string `json:"fqdn,omitempty"`
	// What operating system is installed on the virtual machine.
	GuestOperatingSystem GuestOperatingSystem `json:"guest_operating_system,omitempty"`
	// What time zone is used by the virtual machine (as returned by guest agent).
	GuestTimeZone TimeZone `json:"guest_time_zone,omitempty"`
	// Indicates whether the virtual machine has snapshots with disks in ILLEGAL state.
	HasIllegalImages bool `json:"has_illegal_images,omitempty"`
	// The virtual machine high availability configuration.
	HighAvailability HighAvailability `json:"high_availability,omitempty"`
	ID               string           `json:"id,omitempty"`             // A unique identifier.
	Initialization   Initialization   `json:"initialization,omitempty"` // Reference to the virtual machine’s initialization configuration.
	IO               Io               `json:"io,omitempty"`             // For performance tuning of IO threading.
	Large_icon       Icon             `json:"large___icon,omitempty"`   // Virtual machine’s large icon.
	// Reference to the storage domain this virtual machine/template lease reside on.
	Lease                       StorageDomainLease            `json:"lease,omitempty"`
	Memory                      int                           `json:"memory,omitempty"`                         // The virtual machine’s memory, in bytes.
	MemoryPolicy                MemoryPolicy                  `json:"memory_policy,omitempty"`                  // Reference to virtual machine’s memory management configuration.
	Migration                   MigrationOptions              `json:"migration,omitempty"`                      // Reference to configuration of migration of running virtual machine to another host.
	MigrationDowntime           int                           `json:"migration_downtime,omitempty"`             // Maximum time the virtual machine can be non responsive during its live migration to another host in ms.
	MultiQueuesEnabled          bool                          `json:"multi_queues_enabled,omitempty"`           // If true, each virtual interface will get the optimal number of queues, depending on the available virtual Cpus.
	Name                        string                        `json:"name,omitempty"`                           // A human-readable name in plain text.
	NextRunConfigurationExists  bool                          `json:"next_run_configuration_exists,omitempty"`  // Virtual machine configuration has been changed and requires restart of the virtual machine.
	NumaTuneMode                NumaTuneMode                  `json:"numa_tune_mode,omitempty"`                 // How the NUMA topology is applied.
	Origin                      string                        `json:"origin,omitempty"`                         // The origin of this virtual machine.
	OS                          OperatingSystem               `json:"os,omitempty"`                             // Operating system type installed on the virtual machine.
	Payloads                    []Payload                     `json:"payloads,omitempty"`                       // Optional payloads of the virtual machine, used for ISOs to configure it.
	PlacementPolicy             VmPlacementPolicy             `json:"placement_policy,omitempty"`               // The configuration of the virtual machine’s placement policy.
	RngDevice                   RngDevice                     `json:"rng_device,omitempty"`                     // Random Number Generator device configuration for this virtual machine.
	RunOnce                     bool                          `json:"run_once,omitempty"`                       // If true, the virtual machine has been started using the run once command, meaning it’s configuration might differ from the stored one for the purpose of this single run.
	SerialNumber                SerialNumber                  `json:"serial_number,omitempty"`                  // Virtual machine’s serial number in a cluster.
	SmallIcon                   Icon                          `json:"small_icon,omitempty"`                     // Virtual machine’s small icon.
	SoundcardEnabled            bool                          `json:"soundcard_enabled,omitempty"`              // If true, the sound card is added to the virtual machine.
	SSO                         Sso                           `json:"sso,omitempty"`                            // Reference to the Single Sign On configuration this virtual machine is configured for.
	StartPaused                 bool                          `json:"start_paused,omitempty"`                   // If true, the virtual machine will be initially in 'paused' state after start.
	StartTime                   time.Duration                 `json:"start_time,omitempty"`                     // The date in which the virtual machine was started.
	Stateless                   bool                          `json:"stateless,omitempty"`                      // If true, the virtual machine is stateless - it’s state (disks) are rolled-back after shutdown.
	Status                      VmStatus                      `json:"status,omitempty"`                         // The current status of the virtual machine.
	StatusDetail                string                        `json:"status_detail,omitempty"`                  // Human readable detail of current status.
	StopReason                  string                        `json:"stop_reason,omitempty"`                    // The reason the virtual machine was stopped.
	StopTime                    time.Duration                 `json:"stop_time,omitempty"`                      // The date in which the virtual machine was stopped.
	StorageErrorResumeBehaviour VmStorageErrorResumeBehaviour `json:"storage_error_resume_behaviour,omitempty"` // Determines how the virtual machine will be resumed after storage error.
	TimeZone                    TimeZone                      `json:"time_zone,omitempty"`                      // The virtual machine’s time zone set by oVirt.
	TunnelMigration             bool                          `json:"tunnel_migration,omitempty"`               // If true, the network data transfer will be encrypted during virtual machine live migration.
	Type                        VmType                        `json:"type,omitempty"`                           // Determines whether the virtual machine is optimized for desktop or server.
	USB                         Usb                           `json:"usb,omitempty"`                            // Configuration of USB devices for this virtual machine (count, type).
	UseLatestTemplateVersion    bool                          `json:"use_latest_template_version,omitempty"`    // If true, the virtual machine is reconfigured to the latest version of it’s template when it is started.
	VirtioScsi                  VirtioScsi                    // Reference to VirtIO SCSI configuration.
}

type Bios struct {
	BootMenu BootMenu
	Type     BiosType // Chipset and BIOS type combination.
}

// Represents boot menu configuration for virtual machines and templates.
type BootMenu struct {
	Enabled bool //Whether the boot menu is enabled for this virtual machine (or template), or not.
}

// BiosType enum
//
// Type representing a chipset and a BIOS type combination.
//   - i440fx_sea_bios	-	i440fx chipset with SeaBIOS. For non-x86 architectures this is the only value allowed.
//   - q35_ovmf	-	q35 chipset with OVMF (UEFI) BIOS.
//   - q35_sea_bios	-	q35 chipset with SeaBIOS.
//   - q35_secure_boot	-	q35 chipset with OVMF (UEFI) BIOS with SecureBoot enabled.
type BiosType string

// Representation for serial console device.
type Console struct {
	Enabled bool
}

type Cpu struct {
	Architecture Architecture `json:"architecture,omitempty"`
	Cores        []Core       `json:"cores,omitempty"`
	CPUTune      CpuTune      `json:"cpu_tune,omitempty"`
	Level        int          `json:"level,omitempty"`
	Mode         CpuMode      `json:"mode,omitempty"`
	Name         string       `json:"name,omitempty"`
	Speed        float64      `json:"speed,omitempty"`
	Topology     CpuTopology  `json:"topology,omitempty"`
	Type         string       `json:"type,omitempty"`
}

type Core struct {
	Index  int `json:"index,omitempty"`
	Socket int `json:"socket,omitempty"`
}

type CpuTune struct {
	VcpuPins []VcpuPin `json:"vcpu_pins,omitempty"`
}

type VcpuPin struct {
	CpuSet string `json:"cpu_set,omitempty"`
	Vcpu   int    `json:"vcpu,omitempty"`
}

// Architecture enum of:
//   - ppc64
//   - s390x(IBM S390X CPU architecture.)
//   - undefined
//   - x86_64
type Architecture string

// CpuMode enum of:
//   - custom
//   - host_model
//   - host_passthrough
type CpuMode string

type CpuTopology struct {
	Cores   int `json:"cores,omitempty"`
	Sockets int `json:"sockets,omitempty"`
	Threads int `json:"threads,omitempty"`
}

type Version struct {
	Build int `json:"build,omitempty"`
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	FullVersion string `json:"full_version,omitempty"`
	// A unique identifier.
	ID    string `json:"id,omitempty"`
	Major int    `json:"major,omitempty"`
	Minor int    `json:"minor,omitempty"`
	// A human-readable name in plain text.
	Name     string `json:"name,omitempty"`
	Revision int    `json:"revision,omitempty"`
}

// Properties sent to VDSM to configure various hooks.
type CustomProperty struct {
	// Property name.
	Name string `json:"name,omitempty"`
	// A regular expression defining the available values a custom property can get.
	Regexp string `json:"regexp,omitempty"`
	// Property value.
	Value string `json:"value,omitempty"`
}

// Represents a graphic console configuration.
type Display struct {
	// The IP address of the guest to connect the graphic console client to.
	Address string `json:"address,omitempty"`
	// Indicates if to override the display address per host.
	AllowOverride bool `json:"allow_override,omitempty"`
	// The TLS certificate in case of a TLS connection.
	Certificate Certificate `json:"certificate,omitempty"`
	// Indicates whether a user is able to copy and paste content from an external host into the graphic console.
	CopyPasteEnabled bool `json:"copy_paste_enabled,omitempty"`
	// Returns the action that will take place when the graphic console is disconnected.
	DisconnectAction string `json:"disconnect_action,omitempty"`
	// Indicates if a user is able to drag and drop files from an external host into the graphic console.
	FileTransferEnabled bool `json:"file_transfer_enabled,omitempty"`
	// The keyboard layout to use with this graphic console.
	KeyboardLayout string `json:"keyboard_layout,omitempty"`
	// The number of monitors opened for this graphic console.
	Monitors int `json:"monitors,omitempty"`
	// The port address on the guest to connect the graphic console client to.
	Port int `json:"port,omitempty"`
	// The proxy IP which will be used by the graphic console client to connect to the guest.
	Proxy string `json:"proxy,omitempty"`
	// The secured port address on the guest, in case of using TLS, to connect the graphic console client to.
	SecurePort int `json:"secure_port,omitempty"`
	// Indicates if to use one PCI slot for each monitor or to use a single PCI channel for all multiple monitors.
	SingleQxlPCI bool `json:"single_qxl_pci,omitempty"`
	// Indicates if to use smart card authentication.
	SmartcardEnabled bool `json:"smartcard_enabled,omitempty"`
	// The graphic console protocol type.
	Type DisplayType `json:"type,omitempty"`
}

type Certificate struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	Content string `json:"content,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Subject      string `json:"subject,omitempty"`
}

// Represents an enumeration of the protocol used to connect to the graphic console of the virtual machine.
type DisplayType struct {
	// Display of type SPICE.
	//
	// Display of type SPICE. See https://www.spice-space.org for more details.
	SPICE string `json:"spice,omitempty"`
	// Display of type VNC.
	//
	// Display of type VNC. VNC stands for Virtual Network Computing, and it is a graphical desktop sharing
	// system that uses RFB (Remote Frame Buffer) protocol to remotely control another machine.
	VNC string `json:"vnc,omitempty"`
}

// This type represents a directory service domain.
type Domain struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	User User   `json:"user,omitempty"`
}

// Represents a user in the system.
type User struct {
	// Free text containing comments about this object.
	Comment    string `json:"comment,omitempty"`
	Department string `json:"department,omitempty"`
	// A human-readable description in plain text.
	Description   string `json:"description,omitempty"`
	DomainEntryID string `json:"domain_entry_id,omitempty"`
	Email         string `json:"email,omitempty"`
	// A unique identifier.
	ID       string `json:"id,omitempty"`
	LastName string `json:"last_name,omitempty"`
	LoggedIn bool   `json:"logged_in,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	// Namespace where the user resides.
	Namespace string `json:"namespace,omitempty"`
	Password  string `json:"password,omitempty"`
	// Similar to user_name.
	Principal string `json:"principal,omitempty"`
	// The user’s username.
	UserName string `json:"user_name,omitempty"`
}

// Represents an operating system installed on the virtual machine.
type GuestOperatingSystem struct {
	// The architecture of the operating system, such as x86_64.
	Architecture string `json:"architecture,omitempty"`
	// Code name of the operating system, such as Maipo.
	Codename string `json:"codename,omitempty"`
	// Full name of operating system distribution.
	Distribution string `json:"distribution,omitempty"`
	// Family of operating system, such as Linux.
	Family string `json:"family,omitempty"`
	// Kernel version of the operating system.
	Kernel Kernel `json:"kernel,omitempty"`
	// Version of the installed operating system.
	Version Version `json:"version,omitempty"`
}

type Kernel struct {
	Version Version `json:"version,omitempty"`
}

// Time zone representation.
type TimeZone struct {
	// Name of the time zone.
	Name string `json:"name,omitempty"`
	// UTC offset.
	//
	// Offset from https://en.wikipedia.org/wiki/Coordinated_Universal_Time.
	UTCOffset string `json:"utc_offset,omitempty"`
}

type HighAvailability struct {
	// Define if the virtual machine is considered highly available.
	Enabled bool `json:"enabled,omitempty"`
	// Indicates the priority of the virtual machine inside the run and migration queues.
	Priority int `json:"priority,omitempty"`
}

type Initialization struct {
	Active_directory_ou string `json:"active___directory___ou,omitempty"`
	Authorized_ssh_keys string `json:"authorized___ssh___keys,omitempty"`
	// Deprecated attribute to specify cloud-init configuration.
	CloudInit CloudInit `json:"cloud_init,omitempty"`
	// Attribute specifying the cloud-init protocol to use for formatting the cloud-init network parameters.
	CloudInitNetworkProtocol CloudInitNetworkProtocol `json:"cloud_init_network_protocol,omitempty"`
	Configuration            Configuration            `json:"configuration,omitempty"`
	Custom_script            string                   `json:"custom___script,omitempty"`
	Dns_search               string                   `json:"dns___search,omitempty"`
	Dns_servers              string                   `json:"dns___servers,omitempty"`
	Domain                   string                   `json:"domain,omitempty"`
	HostName                 string                   `json:"host_name,omitempty"`
	InputLocale              string                   `json:"input_locale,omitempty"`
	NicConfiguration         []NicConfiguration       `json:"nic_configuration,omitempty"`
	OrgName                  string                   `json:"org_name,omitempty"`
	RegenerateIDs            bool                     `json:"regenerate_i_ds,omitempty"`
	Regenerate_ssh_keys      bool                     `json:"regenerate___ssh___keys,omitempty"`
	Root_password            string                   `json:"root___password,omitempty"`
	System_locale            string                   `json:"system___locale,omitempty"`
	Timezone                 string                   `json:"timezone,omitempty"`
	UILanguage               string                   `json:"ui_language,omitempty"`
	User_locale              string                   `json:"user___locale,omitempty"`
	User_name                string                   `json:"user___name,omitempty"`
	WindowsLicenseKey        string                   `json:"windows_license_key,omitempty"`
}

// Deprecated type to specify cloud-init configuration.
//
// This type has been deprecated and replaced by alternative attributes inside the Initialization type. See the cloud_init attribute documentation for details.
type CloudInit struct {
	AuthorizedKeys       []AuthorizedKey
	Files                []File
	Host                 Host
	NetworkConfiguration NetworkConfiguration
	RegenerateSshKeys    bool
	Timezone             string
	Users                []User
}

type AuthorizedKey struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
}

type File struct {
	// Free text containing comments about this object.
	Comment string `json:"comment,omitempty"`
	Content string `json:"content,omitempty"`
	// A human-readable description in plain text.
	Description string `json:"description,omitempty"`
	// A unique identifier.
	ID string `json:"id,omitempty"`
	// A human-readable name in plain text.
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Type representing a host.
type Host struct {
	// The host address (FQDN/IP).
	Address string
	// The host auto non uniform memory access (NUMA) status.
	AutoNumaStatus AutoNumaStatus
	// The host certificate.
	Certificate Certificate
	// Free text containing comments about this object.
	Comment string
	// The CPU type of this host.
	CPU Cpu
	// A human-readable description in plain text.
	Description string
	// Specifies whether host device passthrough is enabled on this host.
	DevicePassthrough HostDevicePassthrough
	// Optionally specify the display address of this host explicitly.
	Display Display
	// The host external status.
	ExternalStatus ExternalStatus
	// The host hardware information.
	HardwareInformation HardwareInformation
	// The self-hosted engine status of this host.
	HostedEngine HostedEngine
	// A unique identifier.
	ID string
	// The host iSCSI details.
	ISCSI IscsiDetails
	// The host KDUMP status.
	KdumpStatus KdumpStatus
	// Kernel SamePage Merging (KSM) reduces references to memory pages from multiple identical pages to a single page reference.
	KSM Ksm
	// The host libvirt version.
	LibvirtVersion Version
	// The max scheduling memory on this host in bytes.
	MaxSchedulingMemory int
	// The amount of physical memory on this host in bytes.
	Memory int
	// A human-readable name in plain text.
	Name string
	// Specifies whether a network-related operation, such as 'setup networks', 'sync networks', or 'refresh capabilities', is currently being executed on this host.
	NetworkOperationInProgress bool
	// Specifies whether non uniform memory access (NUMA) is supported on this host.
	NumaSupported bool
	// The operating system on this host.
	OS OperatingSystem
	// Specifies whether we should override firewall definitions.
	OverrideIptables bool
	// The host port.
	Port int
	// The host power management definitions.
	PowerManagement PowerManagement
	// The protocol that the engine uses to communicate with the host.
	Protocol HostProtocol
	// When creating a new host, a root password is required if the password authentication method is chosen, but this is not subsequently included in the representation.
	RootPassword string
	// The host SElinux status.
	SeLinux SeLinux
	// The host storage pool manager (SPM) status and definition.
	SPM Spm
	// The SSH definitions.
	Ssh Ssh
	// The host status.
	Status HostStatus
	// The host status details.
	StatusDetail string
	// The virtual machine summary - how many are active, migrating and total.
	Summary VmSummary
	// Transparent huge page support expands the size of memory pages beyond the standard 4 KiB limit.
	TransparentHugePages TransparentHugePages
	// Indicates if the host contains a full installation of the operating system or a scaled-down version intended only to host virtual machines.
	Type HostType
	// Specifies whether there is an oVirt-related update on this host.
	UpdateAvailable bool
	// The version of VDSM.
	Version Version
	// Specifies the vGPU placement strategy.
	VgpuPlacement VgpuPlacement
}

// AutoNumaStatus enum of:
//   - disable
//   - enable
//   - unknown
type AutoNumaStatus string

type HostDevicePassthrough struct {
	Enabled bool
}

// ExternalStatus enum
//
// Represents an external status. This status is currently used for hosts and storage domains,
// and allows an external system to update status of objects it is aware of.
//   - error - Error status. There is some kind of error in the relevant object.
//   - failure - Failure status. The relevant object is failing.
//   - info - Info status. The relevant object is in OK status,
//     but there is an information available that might be relevant for the administrator.
//   - ok - OK status. The relevant object is working well.
//   - warning - Warning status. The relevant object is working well,
//     but there is some warning that might be relevant for the administrator.
type ExternalStatus string

// Represents hardware information of host.
type HardwareInformation struct {
	// Type of host’s CPU.
	Family string `json:"family,omitempty"`
	// Manufacturer of the host’s machine and hardware vendor.
	Manufacturer string `json:"manufacturer,omitempty"`
	// Host’s product name (for example RHEV Hypervisor).
	ProductName string `json:"product_name,omitempty"`
	// Unique ID for host’s chassis.
	SerialNumber string `json:"serial_number,omitempty"`
	// Supported sources of random number generator.
	SupportedRNGSources []RngSource `json:"supported_rng_sources,omitempty"`
	// Unique ID for each host.
	UUID string `json:"uuid,omitempty"`
	// Unique name for each of the manufacturer.
	Version string `json:"version,omitempty"`
}

// RngSource enum
//
// Representing the random generator backend types.
//   - hwrng - Obtains random data from the /dev/hwrng (usually specialized HW generator) device.
//   - random - Obtains random data from the /dev/random device.
//   - urandom - Obtains random data from the /dev/urandom device.
type RngSource string

type HostedEngine struct {
	Active            bool `json:"active,omitempty"`
	Configured        bool `json:"configured,omitempty"`
	GlobalMaintenance bool `json:"global_maintenance,omitempty"`
	LocalMaintenance  bool `json:"local_maintenance,omitempty"`
	Score             int  `json:"score,omitempty"`
}

type IscsiDetails struct {
	Address         string `json:"address,omitempty"`
	DiskID          string `json:"disk_id,omitempty"`
	Initiator       string `json:"initiator,omitempty"`
	LunMapping      int    `json:"lun_mapping,omitempty"`
	Password        string `json:"password,omitempty"`
	Paths           int    `json:"paths,omitempty"`
	Port            int    `json:"port,omitempty"`
	Portal          string `json:"portal,omitempty"`
	ProductID       string `json:"product_id,omitempty"`
	Serial          string `json:"serial,omitempty"`
	Size            int    `json:"size,omitempty"`
	Status          string `json:"status,omitempty"`
	StorageDomainID string `json:"storage_domain_id,omitempty"`
	Target          string `json:"target,omitempty"`
	Username        string `json:"username,omitempty"`
	VendorID        string `json:"vendor_id,omitempty"`
	VolumeGroupID   string `json:"volume_group_id,omitempty"`
}

// KdumpStatus enum
//   - disabled
//   - enabled
//   - unknown
type KdumpStatus string

type Ksm struct {
	Enabled          bool `json:"enabled,omitempty"`
	MergeAcrossNodes bool `json:"merge_across_nodes,omitempty"`
}

// Information describing the operating system. This is used for both virtual machines and hosts.
type OperatingSystem struct {
	// Configuration of the boot sequence.
	Boot Boot `json:"boot,omitempty"`
	// Custom kernel parameters for start the virtual machine with if Linux operating system is used.
	Cmdline string `json:"cmdline,omitempty"`
	// A custom part of the host kernel command line.
	CustomKernelCmdline string `json:"custom_kernel_cmdline,omitempty"`
	// Path to custom initial ramdisk on ISO storage domain if Linux operating system is used.
	Initrd string `json:"initrd,omitempty"`
	// Path to custom kernel on ISO storage domain if Linux operating system is used.
	Kernel string `json:"kernel,omitempty"`
	// The host kernel command line as reported by a running host.
	ReportedKernelCmdline string `json:"reported_kernel_cmdline,omitempty"`
	// Operating system name in human readable form.
	Type    string  `json:"type,omitempty"`
	Version Version `json:"version,omitempty"`
}

// Configuration of the boot sequence of a virtual machine.
type Boot struct {
	// Ordered list of boot devices.
	// The virtual machine will try to boot from the given boot devices, in the given order.
	Devices []BootDevice `json:"devices,omitempty"`
}

// Represents the kinds of devices that a virtual machine can boot from.
//   - cdrom - Boot from CD-ROM.
//   - hd - Boot from the hard drive.
//   - network - Boot from the network, using PXE.
type BootDevice string

type PowerManagement struct {
	Address            string                `json:"address,omitempty"`              //	The host name or IP address of the host.
	Agents             []Agent               `json:"agents,omitempty"`               //	Specifies fence agent options when multiple fences are used.
	AutomaticPmEnabled bool                  `json:"automatic_pm_enabled,omitempty"` //	Toggles the automated power control of the host in order to save energy.
	Enabled            bool                  `json:"enabled,omitempty"`              //	Indicates whether power management configuration is enabled or disabled.
	KdumpDetection     bool                  `json:"kdump_detection,omitempty"`      //	Toggles whether to determine if kdump is running on the host before it is shut down.
	Options            []Option              `json:"options,omitempty"`              //	Fencing options for the selected type= specified with the option name="" and value="" strings.
	Password           string                `json:"password,omitempty"`             //	A valid, robust password for power management.
	PmProxies          []PmProxy             `json:"pm_proxies,omitempty"`           //	Determines the power management proxy.
	Status             PowerManagementStatus `json:"status,omitempty"`               //	Determines the power status of the host.
	Type               string                `json:"type,omitempty"`                 //	Fencing device code.
	Username           string                `json:"username,omitempty"`             //	A valid user name for power management.
}

// Type representing a fence agent.
type Agent struct {
	Address        string   `json:"address,omitempty"`         //	Fence agent address.
	Comment        string   `json:"comment,omitempty"`         //	Free text containing comments about this object.
	Concurrent     bool     `json:"concurrent,omitempty"`      //	Specifies whether the agent should be used concurrently or sequentially.
	Description    string   `json:"description,omitempty"`     //	A human-readable description in plain text.
	EncryptOptions bool     `json:"encrypt_options,omitempty"` //	Specifies whether the options should be encrypted.
	ID             string   `json:"id,omitempty"`              //	A unique identifier.
	Name           string   `json:"name,omitempty"`            //	A human-readable name in plain text.
	Options        []Option `json:"options,omitempty"`         //	Fence agent options (comma-delimited list of key-value pairs).
	Order          int      `json:"order,omitempty"`           //	The order of this agent if used with other agents.
	Password       string   `json:"password,omitempty"`        //	Fence agent password.
	Port           int      `json:"port,omitempty"`            //	Fence agent port.
	Type           string   `json:"type,omitempty"`            //	Fence agent type.
	Username       string   `json:"username,omitempty"`        //	Fence agent user name.
}

type Option struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type PmProxy struct {
	Type PmProxyType `json:"type,omitempty"`
}

// PmProxyType enum
//   - cluster - The fence proxy is selected from the same cluster as the fenced host.
//   - dc - The fence proxy is selected from the same data center as the fenced host.
//   - other_dc - The fence proxy is selected from a different data center than the fenced host.
type PmProxyType string

// PowerManagementStatus enum
//   - off - Host is OFF.
//   - on - Host is ON.
//   - unknown - Unknown status.
type PowerManagementStatus string

// HostProtocol enum
//
// The protocol used by the engine to communicate with a host.
//   - stomp -	JSON-RPC protocol on top of STOMP.
//   - xml -	XML-RPC protocol.
//
// Since version 4.1 of the engine the protocol is always set to stomp since xml was removed.
type HostProtocol string

// Represents SELinux in the system.
type SeLinux struct {
	Mode SeLinuxMode `json:"mode,omitempty"` // SELinux current mode.
}

// SeLinuxMode enum
//
// Represents an SELinux enforcement mode.
//   - disabled	- SELinux is disabled in the kernel.
//   - enforcing	- SELinux is running and enforcing permissions.
//   - permissive	- SELinux is running and logging but not enforcing permissions.
type SeLinuxMode string

type Spm struct {
	Priority int       `json:"priority,omitempty"`
	Status   SpmStatus `json:"status,omitempty"`
}

// SpmStatus enum
//   - contending
//   - none
//   - spm
type SpmStatus string

type Ssh struct {
	AuthenticationMethod SshAuthenticationMethod `json:"authentication_method,omitempty"`
	Comment              string                  `json:"comment,omitempty"`     //	Free text containing comments about this object.
	Description          string                  `json:"description,omitempty"` //	A human-readable description in plain text.
	Fingerprint          string                  `json:"fingerprint,omitempty"`
	ID                   string                  `json:"id,omitempty"`   //	A unique identifier.
	Name                 string                  `json:"name,omitempty"` //	A human-readable name in plain text.
	Port                 int                     `json:"port,omitempty"`
	User                 User                    `json:"user,omitempty"`
}

// SshAuthenticationMethod enum
//   - password
//   - publickey
type SshAuthenticationMethod string

// HostStatus enum
//
// Type representing a host status.
//   - connecting	 - 	The engine cannot communicate with the host for a specific threshold so it is now trying to connect before going through fencing.
//   - down	 - 	The host is down.
//   - error	 - 	The host is in error status.
//   - initializing	 - 	The host is initializing.
//   - install_failed	 - 	The host installation failed.
//   - installing	 - 	The host is being installed.
//   - installing_os	 - 	The host operating system is now installing.
//   - kdumping	 - 	The host kernel has crashed and it is now going through memory dumping.
//   - maintenance	 - 	The host is in maintenance status.
//   - non_operational	 - 	The host is non operational.
//   - non_responsive	 - 	The host is not responsive.
//   - pending_approval	 - 	The host is pending administrator approval.
//   - preparing_for_maintenance	 - 	The host is preparing for maintenance.
//   - reboot	 - 	The host is being rebooted.
//   - unassigned	 - 	The host is in activation process.
//   - up	 - 	The host is up.
type HostStatus string

// Type containing information related to virtual machines on a particular host.
type VmSummary struct {
	Active    int `json:"active,omitempty"`    //	The number of virtual machines active on the host.
	Migrating int `json:"migrating,omitempty"` //	The number of virtual machines migrating to or from the host.
	Total     int `json:"total,omitempty"`     //	The number of virtual machines present on the host.
}

// Type representing a transparent huge pages (THP) support.
type TransparentHugePages struct {
	Enabled bool `json:"enabled,omitempty"` // Enable THP support.
}

// HostType enum
//
// This enumerated type is used to determine which type of operating system is used by the host.
//   - ovirt_node	 - 	The host contains Red Hat Virtualization Host (RHVH):
//     a new implementation of Red Hat Enterprise Virtualization Hypervisor (RHEV-H)
//     which uses the same installer as Red Hat Enterprise Linux, CentOS, or Fedora.
//   - rhel	 - 	The host contains a full Red Hat Enterprise Linux, CentOS, or Fedora installation.
//   - rhev_h	 - 	The host contains Red Hat Enterprise Virtualization Hypervisor (RHEV-H),
//     a small-scaled version of Red Hat Enterprise Linux, CentOS, or Fedora, used solely to host virtual machines.
type HostType string

// VgpuPlacement enum
//
// The vGPU placement strategy.
// It can either put vGPUs on the first available physical cards, or spread them over multiple physical cards.
//   - consolidated - Use consolidated placement. Each vGPU is placed on the first physical card with available space.
//     This is the default placement, utilizing all available space on the physical cards.
//   - separated - Use separated placement. Each vGPU is placed on a separate physical card, if possible.
//     This can be useful for improving vGPU performance.
type VgpuPlacement string

type NetworkConfiguration struct {
	DNS  Dns
	Nics []Nic
}

// Represents the DNS resolver configuration.
type Dns struct {
	SearchDomains []Host `json:"search_domains,omitempty"` //	Array of hosts serving as search domains.
	Servers       []Host `json:"servers,omitempty"`        //	Array of hosts serving as DNS servers.
}

// Represents a virtual machine NIC.
type Nic struct {
	BootProtocol BootProtocol `json:"boot_protocol,omitempty"` //	Defines how an IP address is assigned to the NIC.
	Comment      string       `json:"comment,omitempty"`       //	Free text containing comments about this object.
	Description  string       `json:"description,omitempty"`   //	A human-readable description in plain text.
	ID           string       `json:"id,omitempty"`            //	A unique identifier.
	Interface    NicInterface `json:"interface,omitempty"`     //	The type of driver used for the NIC.
	Linked       bool         `json:"linked,omitempty"`        //	Defines if the NIC is linked to the virtual machine.
	MAC          Mac          `json:"mac,omitempty"`           //	The MAC address of the interface.
	Name         string       `json:"name,omitempty"`          //	A human-readable name in plain text.
	OnBoot       bool         `json:"on_boot,omitempty"`       //	Defines if the network interface should be activated upon operation system startup.
	Plugged      bool         `json:"plugged,omitempty"`       //	Defines if the NIC is plugged in to the virtual machine.
}

// BootProtocol enum
//
// Defines the options of the IP address assignment method to a NIC.
//   - autoconf	 - 	Stateless address auto-configuration.
//   - dhcp	 - 	Dynamic host configuration protocol.
//   - none	 - 	No address configuration.
//   - poly_dhcp_autoconf	 - 	DHCP alongside Stateless address auto-configuration (SLAAC).
//   - static	 - 	Statically-defined address, mask and gateway.
type BootProtocol string

// NicInterface enum
//
// Defines the options for an emulated virtual network interface device model.
//   - e1000	 - 	e1000.
//   - pci_passthrough	 - 	PCI Passthrough.
//   - rtl8139	 - 	rtl8139.
//   - rtl8139_virtio	 - 	Dual mode rtl8139, VirtIO.
//   - spapr_vlan	 - 	sPAPR VLAN.
//   - virtio	 - 	VirtIO.
type NicInterface string

// Represents a MAC address of a virtual network interface.
type Mac struct {
	Address string // MAC address.
}

// CloudInitNetworkProtocol enum
//
// Defines the values for the cloud-init protocol. This protocol decides how the cloud-init network
// parameters are formatted before being passed to the virtual machine in order to be processed by cloud-init.
// Protocols supported are cloud-init version dependent. For more information, see http://cloudinit.readthedocs.io/en/latest/topics/network-config.html#network-configuration-sources
//   - eni - Legacy protocol. Does not support IPv6. For more information, see Network Configuration ENI (Legacy)
//   - openstack_metadata - Successor of the ENI protocol, with support for IPv6 and more. This is the default value. For more information, see http://specs.openstack.org/openstack/nova-specs/specs/liberty/implemented/metadata-service-network-info
type CloudInitNetworkProtocol string

type Configuration struct {
	Data string            `json:"data,omitempty"` //	The document describing the virtual machine.
	Type ConfigurationType `json:"type,omitempty"`
}

// ConfigurationType enum
//
// Configuration format types.
//   - ova - ConfigurationType of type standard OVF.
//     The provided virtual machine configuration conforms with the Open Virtualization Format (OVF) standard.
//     This value should be used for an OVF configuration that is extracted from an Open Virtual Appliance (OVA)
//     that was generated by oVirt or by other vendors. See https://www.dmtf.org/standards/ovf for the OVF specification.
//   - ovf - ConfigurationType of type oVirt-compatible OVF.
//     The provided virtual machine configuration conforms with the oVirt-compatible form of the Open
//     Virtualization Format (OVF). Note that the oVirt-compatible form of the OVF may differ from the OVF
//     standard that is used by other vendors. This value should be used for an OVF configuration that is taken
//     from a storage domain.
type ConfigurationType string

// The type describes the configuration of a virtual network interface.
type NicConfiguration struct {
	BootProtocol     BootProtocol `json:"boot_protocol,omitempty"`      //	IPv4 boot protocol.
	IP               Ip           `json:"ip,omitempty"`                 //	IPv4 address details.
	IPv6             Ip           `json:"ipv6,omitempty"`               //	IPv6 address details.
	IPv6BootProtocol BootProtocol `json:"ipv6_boot_protocol,omitempty"` //	IPv6 boot protocol.
	Name             string       `json:"name,omitempty"`               //	Network interface name.
	OnBoot           bool         `json:"on_boot,omitempty"`            //	Specifies whether the network interface should be activated on the virtual machine guest operating system boot.
}

// Represents the IP configuration of a network interface.
type Ip struct {
	Address string    `json:"address,omitempty"` //	The text representation of the IP address.
	Gateway string    `json:"gateway,omitempty"` //	The address of the default gateway.
	Netmask string    `json:"netmask,omitempty"` //	The network mask.
	Version IpVersion `json:"version,omitempty"` //	The version of the IP protocol.
}

// IpVersion enum
//
// Defines the values for the IP protocol version.
//   - v4 - IPv4.
//   - v6 - IPv6.
type IpVersion string

type Io struct {
	Threads int `json:"threads,omitempty"`
}

// Icon of virtual machine or template.
type Icon struct {
	Comment     string `json:"comment,omitempty"`     //	Free text containing comments about this object.
	Data        string `json:"data,omitempty"`        //	Base64 encode content of the icon file.
	Description string `json:"description,omitempty"` //	A human-readable description in plain text.
	ID          string `json:"id,omitempty"`          //	A unique identifier.
	// Format of icon file.
	//
	// One of:
	//  - image/jpeg
	//  - image/png
	//  - image/gif
	MediaType string `json:"media_type,omitempty"`
	Name      string `json:"name,omitempty"` //	A human-readable name in plain text.
}

// Represents a lease residing on a storage domain.
//
// A lease is a Sanlock http://www.ovirt.org/develop/developer-guide/vdsm/sanlock
// resource residing on a special volume on the storage domain,
// this Sanlock resource is used to provide storage base locking.
type StorageDomainLease struct {
	StorageDomain StorageDomain // Reference to the storage domain on which the lock resides on.
}

// Storage domain.
type StorageDomain struct {
	Available                  int                 `json:"available,omitempty"`
	Backup                     bool                `json:"backup,omitempty"`     //	This attribute indicates whether a data storage domain is used as backup domain or not.
	BlockSize                  int                 `json:"block_size,omitempty"` //	Specifies block size in bytes for a storage domain.
	Comment                    string              `json:"comment,omitempty"`    //	Free text containing comments about this object.
	Committed                  int                 `json:"committed,omitempty"`
	CriticalSpaceActionBlocker int                 `json:"critical_space_action_blocker,omitempty"`
	Description                string              `json:"description,omitempty"`          //	A human-readable description in plain text.
	DiscardAfterDelete         bool                `json:"discard_after_delete,omitempty"` //	Indicates whether disks' blocks on block storage domains will be discarded right before they are deleted.
	ExternalStatus             ExternalStatus      `json:"external_status,omitempty"`
	ID                         string              `json:"id,omitempty"` //	A unique identifier.
	Import                     bool                `json:"import,omitempty"`
	Master                     bool                `json:"master,omitempty"`
	Name                       string              `json:"name,omitempty"` //	A human-readable name in plain text.
	Status                     StorageDomainStatus `json:"status,omitempty"`
	Storage                    HostStorage         `json:"storage,omitempty"`
	StorageFormat              StorageFormat       `json:"storage_format,omitempty"`
	SupportsDiscard            bool                `json:"supports_discard,omitempty"`             //	Indicates whether a block storage domain supports discard operations.
	SupportsDiscardZeroesData  bool                `json:"supports_discard_zeroes_data,omitempty"` //	Indicates whether a block storage domain supports the property that discard zeroes the data.
	Type                       StorageDomainType   `json:"type,omitempty"`
	Used                       int                 `json:"used,omitempty"`
	WarningLowSpaceIndicator   int                 `json:"warning_low_space_indicator,omitempty"`
	WipeAfterDelete            bool                `json:"wipe_after_delete,omitempty"` //	Serves as the default value of wipe_after_delete for disks on this storage domain.
}

// StorageDomainStatus enum
//   - activating
//   - active
//   - detaching
//   - inactive
//   - locked
//   - maintenance
//   - mixed
//   - preparing_for_maintenance
//   - unattached
//   - unknown
type StorageDomainStatus string

type HostStorage struct {
	Address                string        `json:"address,omitempty"`
	Comment                string        `json:"comment,omitempty"`                  //	Free text containing comments about this object.
	Description            string        `json:"description,omitempty"`              //	A human-readable description in plain text.
	DriverOptions          []Property    `json:"driver_options,omitempty"`           //	The options to be passed when creating a storage domain using a cinder driver.
	DriverSensitiveOptions []Property    `json:"driver_sensitive_options,omitempty"` //	Parameters containing sensitive information, to be passed when creating a storage domain using a cinder driver.
	ID                     string        `json:"id,omitempty"`                       //	A unique identifier.
	LogicalUnits           []LogicalUnit `json:"logical_units,omitempty"`
	MountOptions           string        `json:"mount_options,omitempty"`
	Mame                   string        `json:"mame,omitempty"`        //	A human-readable name in plain text.
	NfsRetrans             int           `json:"nfs_retrans,omitempty"` //	The number of times to retry a request before attempting further recovery actions.
	NfsTimeo               int           `json:"nfs_timeo,omitempty"`   //	The time in tenths of a second to wait for a response before retrying NFS requests.
	NfsVersion             NfsVersion    `json:"nfs_version,omitempty"`
	OverrideLuns           bool          `json:"override_luns,omitempty"`
	Password               string        `json:"password,omitempty"`
	Path                   string        `json:"path,omitempty"`
	Port                   int           `json:"port,omitempty"`
	Portal                 string        `json:"portal,omitempty"`
	Target                 string        `json:"target,omitempty"`
	Type                   StorageType   `json:"type,omitempty"`
	Username               string        `json:"username,omitempty"`
	VfsType                string        `json:"vfs_type,omitempty"`
	VolumeGroup            VolumeGroup   `json:"volume_group,omitempty"`
}

type Property struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type LogicalUnit struct {
	Address           string    `json:"address,omitempty"`
	DiscardMaxSize    int       `json:"discard_max_size,omitempty"`    //	The maximum number of bytes that can be discarded by the logical unit’s underlying storage in a single operation.
	DiscardZeroesData bool      `json:"discard_zeroes_data,omitempty"` //	True, if previously discarded blocks in the logical unit’s underlying storage are read back as zeros.
	DiskID            string    `json:"disk_id,omitempty"`
	ID                string    `json:"id,omitempty"`
	LunMapping        int       `json:"lun_mapping,omitempty"`
	Password          string    `json:"password,omitempty"`
	Paths             int       `json:"paths,omitempty"`
	Port              int       `json:"port,omitempty"`
	Portal            string    `json:"portal,omitempty"`
	ProductID         string    `json:"product_id,omitempty"`
	Serial            string    `json:"serial,omitempty"`
	Size              int       `json:"size,omitempty"`
	Status            LunStatus `json:"status,omitempty"`
	StorageDomainID   string    `json:"storage_domain_id,omitempty"`
	Target            string    `json:"target,omitempty"`
	Username          string    `json:"username,omitempty"`
	VendorID          string    `json:"vendor_id,omitempty"`
	VolumeGroupID     string    `json:"volume_group_id,omitempty"`
}

// LunStatus enum
//   - free
//   - unusable
//   - used
type LunStatus string

// NfsVersion enum
//   - auto
//   - v3
//   - v4
//   - v4_0	-	NFS 4.0.
//   - v4_1
//   - v4_2	- 	NFS 4.2.
type NfsVersion string

// StorageType enum
//
// Type representing a storage domain type.
//   - cinder	-	Cinder storage domain.
//   - fcp	-	Fibre-Channel storage domain.
//   - glance	-	Glance storage domain.
//   - glusterfs	-	Gluster-FS storage domain.
//   - iscsi	-	iSCSI storage domain.
//   - localfs	-	Storage domain on Local storage.
//   - managed_block_storage	-	Managed block storage domain. A storage domain managed using cinderlib.
//   - nfs	-	NFS storage domain.
//   - posixfs	-	POSIX-FS storage domain.
type StorageType string

type VolumeGroup struct {
	ID           string
	LogicalUnits []LogicalUnit
	Name         string
}

// StorageFormat enum
//
// Type which represents a format of storage domain.
//   - v1	-	Version 1 of the storage domain format is applicable to NFS, iSCSI and FC storage domains.
//   - v2	-	Version 2 of the storage domain format is applicable to iSCSI and FC storage domains.
//   - v3	-	Version 3 of the storage domain format is applicable to NFS, POSIX, iSCSI and FC storage domains.
//   - v4	-	Version 4 of the storage domain format.
//   - v5	-	Version 5 of the storage domain format is applicable to NFS, POSIX, and Gluster storage domains.
type StorageFormat string

// StorageDomainType enum
//
// Indicates the kind of data managed by a storage domain.
//   - data	-	Data domains are used to store the disks and snapshots of the virtual machines and templates in the system.
//   - export	-	Export domains are temporary storage repositories used to copy and move virtual machines and templates between data centers and oVirt environments.
//   - image	-	Image domain store images that can be imported into from an external system.
//   - iso	-	ISO domains store ISO files (or logical CDs) used to install and boot operating systems and applications for the virtual machines.
//   - managed_block_storage	-	Managed block storage domains are created on block storage devices.
//   - volume	-	Volume domains store logical volumes that can be used as disks for virtual machines.
type StorageDomainType string

// Logical grouping of memory-related properties of virtual machine-like entities.
type MemoryPolicy struct {
	Ballooning           bool                 `json:"ballooning,omitempty"`
	Guaranteed           int                  `json:"guaranteed,omitempty"` //	The amount of memory, in bytes, that is guaranteed to not be drained by the balloon mechanism.
	Max                  int                  `json:"max,omitempty"`        //	Maximum virtual machine memory, in bytes.
	OverCommit           MemoryOverCommit     `json:"over_commit,omitempty"`
	TransparentHugePages TransparentHugePages `json:"transparent_huge_pages,omitempty"`
}

type MemoryOverCommit struct {
	Percent int `json:"percent,omitempty"`
}

// The type for migration options.
type MigrationOptions struct {
	AutoConverge Inheritablebool    `json:"auto_converge,omitempty"`
	Bandwidth    MigrationBandwidth `json:"bandwidth,omitempty"` //	The bandwidth that is allowed to be used by the migration.
	Compressed   Inheritablebool    `json:"compressed,omitempty"`
	Policy       MigrationPolicy    `json:"policy,omitempty"` //	A reference to the migration policy, as defined using engine-config.
}

// Inheritablebool enum
//
// Enum representing the bool value that can be either set, or inherited from a higher level.
// The inheritance order is virtual machine → cluster → engine-config.
//   - false	-	Set the value to false on this level.
//   - inherit	-	Inherit the value from higher level.
//   - true	-	Set the value to true on this level.
type Inheritablebool string

// Defines the bandwidth used by migration.
type MigrationBandwidth struct {
	AssignmentMethod MigrationBandwidthAssignmentMethod //	The method used to assign the bandwidth.
	CustomValue      int                                //	Custom bandwidth in Mbps. Will be applied only if the assignmentMethod attribute is custom.
}

// MigrationBandwidthAssignmentMethod enum
//
// Defines how the migration bandwidth is assigned.
//   - auto	-	Takes the bandwidth from the Quality of Service if the Quality of Service is defined.
//     If the Quality of Service is not defined the bandwidth is taken from the detected link speed being used.
//     If nothing is detected, bandwidth falls back to the hypervisor_default value.
//   - custom	-	Custom defined bandwidth in Mbit/s.
//   - hypervisor_default	-	Takes the value as configured on the hypervisor.
type MigrationBandwidthAssignmentMethod string

// A policy describing how the migration is treated, such as convergence or how many parallel migrations are allowed.
type MigrationPolicy struct {
	Comment     string `json:"comment,omitempty"`     //	Free text containing comments about this object.
	Description string `json:"description,omitempty"` //	A human-readable description in plain text.
	ID          string `json:"id,omitempty"`          //	A unique identifier.
	Name        string `json:"name,omitempty"`        //	A human-readable name in plain text.
}

// NumaTuneMode enum
//
//   - interleave
//   - preferred
//   - strict
type NumaTuneMode string

type Payload struct {
	Files    []File
	Type     VmDeviceType
	VolumeID string
}

// VmDeviceType enum
//
//   - cdrom
//   - floppy
type VmDeviceType string

type VmPlacementPolicy struct {
	Affinity VmAffinity `json:"affinity,omitempty"`
	Hosts    []Host     `json:"hosts,omitempty"`
}

// VmAffinity enum
//
//   - migratable
//   - pinned
//   - user_migratable
type VmAffinity string

// Random number generator (RNG) device model.
type RngDevice struct {
	Rate   Rate      `json:"rate,omitempty"`   //	Determines maximum speed of consumption of bytes from random number generator device.
	Source RngSource `json:"source,omitempty"` //	Backend of the random number generator device.
}

// Determines maximum speed of consumption of bytes from random number generator device.
type Rate struct {
	Bytes  int `json:"bytes,omitempty"`  //	Number of bytes allowed to consume per period.
	Period int `json:"period,omitempty"` //	Duration of one period in milliseconds.
}

type SerialNumber struct {
	Policy SerialNumberPolicy `json:"policy,omitempty"`
	Value  string             `json:"value,omitempty"`
}

// SerialNumberPolicy enum
//
//   - custom
//   - host
//   - vm
type SerialNumberPolicy string

type Sso struct {
	Method []Method
}

type Method struct {
	ID SsoMethod
}

// SsoMethod enum
//
//   - guest_agent
type SsoMethod string

// VmStatus enum
//
// Type represeting a status of a virtual machine.
//   - down	-	This status indicates that the virtual machine process is not running.
//   - image_locked	-	This status indicates that the virtual machine process is not running and there
//     is some operation on the disks of the virtual machine that prevents it from being started.
//   - migrating	-	This status indicates that the virtual machine process is running and the virtual machine
//     is being migrated from one host to another.
//   - not_responding	-	This status indicates that the hypervisor detected that the virtual machine
//     is not responding.
//   - paused	-	This status indicates that the virtual machine process is running and the virtual machine is paused.
//   - powering_down	-	This status indicates that the virtual machine process is running and it is about to
//     stop running.
//   - powering_up	-	This status indicates that the virtual machine process is running and the guest operating
//     system is being loaded.
//   - reboot_in_progress	-	This status indicates that the virtual machine process is running and the guest
//     operating system is being rebooted.
//   - restoring_state	-	This status indicates that the virtual machine process is about to run and the virtual
//     machine is going to awake from hibernation.
//   - saving_state	-	This status indicates that the virtual machine process is running and the virtual
//     machine is being hibernated.
//   - suspended	-	This status indicates that the virtual machine process is not running and a running
//     state of the virtual machine was saved.
//   - unassigned	-	This status is set when an invalid status is received.
//   - unknown	-	This status indicates that the system failed to determine the status of the virtual machine.
//   - up	-	This status indicates that the virtual machine process is running and the guest operating
//     system is loaded.
//   - wait_for_launch	-	This status indicates that the virtual machine process is about to run.
type VmStatus string

// VmStorageErrorResumeBehaviour enum
//
// If the storage, on which this virtual machine has some disks gets unresponsive, the virtual machine gets paused.
// This are the possible options, what should happen with the virtual machine in the moment
// the storage gets available again.
type VmStorageErrorResumeBehaviour string

// VmType enum
//
// Type representing what the virtual machine is optimized for.
//   - desktop	-	The virtual machine is intended to be used as a desktop.
//   - high_performance	-	The virtual machine is intended to be used as a high performance virtual machine.
//   - server	-	The virtual machine is intended to be used as a server.
type VmType string

// Configuration of the USB device of a virtual machine.
type Usb struct {
	Enabled bool    `json:"enabled,omitempty"` //	Determines whether the USB device should be included or not.
	Type    UsbType `json:"type,omitempty"`    //	USB type, currently only native is supported.
}

// UsbType enum
//
// Type of USB device redirection.
//   - legacy	-	Legacy USB redirection. This USB type has been deprecated since version 3.6 of the engine,
//     and has been completely removed in version 4.1. It is preserved only to avoid syntax errors in existing scripts.
//     If it is used it will be automatically replaced by native.
//   - native	-	Native USB redirection. Native USB redirection allows KVM/SPICE USB redirection for Linux and
//     Windows virtual machines. Virtual (guest) machines require no guest-installed agents or drivers for native USB.
//     On Linux clients, all packages required for USB redirection are provided by the virt-viewer package.
//     On Windows clients, you must also install the usbdk package.
type UsbType string

type VirtioScsi struct {
	Enabled bool `json:"enabled,omitempty"` // Enable Virtio SCSI support.
}

// Represents a virtual disk device.
type Disk struct {
	Active                bool //	Indicates if the disk is visible to the virtual machine.
	ActualSize            int  //	The actual size of the disk, in bytes.
	Alias                 string
	Backup                DiskBackup      //	The backup behavior supported by the disk.
	Bootable              bool            //	Indicates if the disk is marked as bootable.
	Comment               string          //	Free text containing comments about this object.
	ContentType           DiskContentType //	Indicates the actual content residing on the disk.
	Description           string          //	A human-readable description in plain text.
	Format                DiskFormat      //	The underlying storage format.
	ID                    string          //	A unique identifier.
	ImageID               string
	InitialSize           int           //	The initial size of a sparse image disk created on block storage, in bytes.
	Interface             DiskInterface //	The type of interface driver used to connect the disk device to the virtual machine.
	LogicalName           string
	LunStorage            HostStorage
	Name                  string        //	A human-readable name in plain text.
	PropagateErrors       bool          //	Indicates if disk errors should cause virtual machine to be paused or if disk errors should be propagated to the the guest operating system instead.
	ProvisionedSize       int           //	The virtual size of the disk, in bytes.
	QcowVersion           QcowVersion   //	The underlying QCOW version of a QCOW volume.
	ReadOnly              bool          //	Indicates if the disk is in read-only mode.
	Sgio                  ScsiGenericIO //	Indicates whether SCSI passthrough is enable and its policy.
	Shareable             bool          //	Indicates if the disk can be attached to multiple virtual machines.
	Sparse                bool          //	Indicates if the physical storage for the disk should not be preallocated.
	Status                DiskStatus    //	The status of the disk device.
	StorageType           DiskStorageType
	TotalSize             int //	The total size of the disk including all of its snapshots, in bytes.
	Uses_scsi_reservation bool
	WipeAfterDelete       bool //	"Indicates if the disk’s blocks will be read back as zeros after it is deleted: - On block storage, the disk will be zeroed and only then deleted."
}

// DiskBackup enum
//
// Represents an enumeration of the backup mechanism that is enabled on the disk.
//   - incremental	-	Incremental backup support.
//   - none	-	No backup support.
type DiskBackup string

// DiskContentType enum
//
// The actual content residing on the disk.
//   - data	-	The disk contains data.
//   - hosted_engine	-	The disk contains the Hosted Engine VM disk.
//   - hosted_engine_configuration	-	The disk contains the Hosted Engine configuration disk.
//   - hosted_engine_metadata	-	The disk contains the Hosted Engine metadata disk.
//   - hosted_engine_sanlock	-	The disk contains the Hosted Engine Sanlock disk.
//   - iso	-	The disk contains an ISO image to be used a CDROM device.
//   - memory_dump_volume	-	The disk contains a memory dump from a live snapshot.
//   - memory_metadata_volume	-	The disk contains memory metadata from a live snapshot.
//   - ovf_store	-	The disk is an OVF store.
type DiskContentType string

// DiskFormat enum
//
// The underlying storage format of disks.
//   - cow	-	The Copy On Write format allows snapshots, with a small performance overhead.
//   - raw	-	The raw format does not allow snapshots, but offers improved performance.
type DiskFormat string

// DiskInterface enum
//
// The underlying storage interface of disks communication with controller.
//   - ide	-	Legacy controller device. Works with almost all guest operating systems, so it is good for compatibility.
//     Performance is lower than with the other alternatives.
//   - sata	-	SATA controller device.
//   - spapr_vscsi	-	Para-virtualized device supported by the IBM pSeries family of machines, using the SCSI protocol.
//   - virtio	-	Virtualization interface where just the guest’s device driver knows it is running in
//     a virtual environment. Enables guests to get high performance disk operations.
//   - virtio_scsi	-	Para-virtualized SCSI controller device. Fast interface with the guest via direct
//     physical storage device address, using the SCSI protocol.
type DiskInterface string

// QcowVersion enum
//
// The QCOW version specifies to the qemu which qemu version the volume supports.
//
// This field can be updated using the update API and will be reported only for QCOW volumes, it is determined
// by the storage domain’s version which the disk is created on. Storage domains with version lower than V4 support
// QCOW2 version 2 volumes, while V4 storage domains also support QCOW2 version 3. For more information about
// features of the different QCOW versions, see http://wiki.qemu.org/Features/Qcow3
//   - qcow2_v2 - The Copy On Write default compatibility version It means that every QEMU can use it.
//   - qcow2_v3 - The Copy On Write compatibility version which was introduced in QEMU 1.1 It means that the new
//     format is in use.
type QcowVersion string

// ScsiGenericIO enum
//
// When a direct LUN disk is using SCSI passthrough the privileged I/O policy is determined by this enum.
//   - disabled	-	Disable SCSI passthrough.
//   - filtered	-	Disallow privileged SCSI I/O.
//   - unfiltered	-	Allow privileged SCSI I/O.
type ScsiGenericIO string

// DiskStatus enum
//
// Current status representation for disk.
//   - illegal	-	Disk cannot be accessed by the virtual machine, and the user needs to take action to resolve the issue.
//   - locked	-	The disk is being used by the system, therefore it cannot be accessed by virtual machines at this point. This is usually a temporary status, until the disk is freed.
//   - ok	-	The disk status is normal and can be accessed by the virtual machine.
type DiskStatus string

// DiskStorageType enum
//
//   - cinder
//   - image
//   - lun
//   - managed_block_storage -	A storage type, used for a storage domain that was created using a cinderlib driver.
type DiskStorageType string

// Describes how a disk is attached to a virtual machine.
type DiskAttachment struct {
	//	Defines whether the disk is active in the virtual machine it’s attached to.
	// A disk attached to a virtual machine in an active status is connected to the virtual machine at run time
	// and can be used.
	Active      bool   `json:"active,omitempty"`
	Bootable    bool   `json:"bootable,omitempty"`    //	Defines whether the disk is bootable.
	Comment     string `json:"comment,omitempty"`     //	Free text containing comments about this object.
	Description string `json:"description,omitempty"` //	A human-readable description in plain text.
	ID          string `json:"id,omitempty"`          //	A unique identifier.
	//	The type of interface driver used to connect the disk device to the virtual machine.
	Interface DiskInterface `json:"interface,omitempty"`
	//	The logical name of the virtual machine’s disk, as seen from inside the virtual machine.
	//
	//	 The logical name of a disk is reported only when the guest agent is installed and running inside the virtual machine.
	//	 If the guest operating system is Windows, the logical name will be reported as \\.\PHYSICALDRIVE0.
	LogicalName string `json:"logical_name,omitempty"`
	Name        string `json:"name,omitempty"` //	A human-readable name in plain text.
	//	Defines whether the virtual machine passes discard commands to the storage.
	PassDiscard bool `json:"pass_discard,omitempty"`
	//	Indicates whether the disk is connected to the virtual machine as read only.
	//	When adding a new disk attachment the default value is false.
	ReadOnly bool `json:"read_only,omitempty"`
	// Defines whether SCSI reservation is enabled for this disk. Virtual machines with VIRTIO-SCSI passthrough
	// enabled can set persistent SCSI reservations on disks. If they set persistent SCSI reservations,
	// those virtual machines cannot be migrated to a different host because they would lose access to the disk,
	// because SCSI reservations are specific to SCSI initiators, and therefore hosts. This scenario cannot be
	// automatically detected. To avoid migrating these virtual machines, the user can set this attribute to true,
	// to indicate the virtual machine is using SCSI reservations.
	UsesScsiReservation bool `json:"uses_scsi_reservation,omitempty"`
}

// The type that represents a virtual machine template. Templates allow for a rapid instantiation of
// virtual machines with common configuration and disk states.
type Template struct {
	Bios                        Bios                          `json:"bios,omitempty"`                         //	Reference to virtual machine’s BIOS configuration.
	Comment                     string                        `json:"comment,omitempty"`                      //	Free text containing comments about this object.
	Console                     Console                       `json:"console,omitempty"`                      //	Console configured for this virtual machine.
	Cpu                         Cpu                           `json:"cpu,omitempty"`                          //	The configuration of the virtual machine CPU.
	CpuShares                   int                           `json:"cpu_shares,omitempty"`                   //
	CreationTime                time.Duration                 `json:"creation_time,omitempty"`                //	The virtual machine creation date.
	CustomCompatibilityVersion  Version                       `json:"custom_compatibility_version,omitempty"` //	Virtual machine custom compatibility version.
	CustomCpuModel              string                        `json:"custom_cpu_model,omitempty"`
	CustomEmulatedMachine       string                        `json:"custom_emulated_machine,omitempty"`
	CustomProperties            []CustomProperty              `json:"custom_properties,omitempty"`              //	Properties sent to VDSM to configure various hooks.
	DeleteProtected             bool                          `json:"delete_protected,omitempty"`               //	If true, the virtual machine cannot be deleted.
	Description                 string                        `json:"description,omitempty"`                    //	A human-readable description in plain text.
	Display                     Display                       `json:"display,omitempty"`                        //	The virtual machine display configuration.
	Domain                      Domain                        `json:"domain,omitempty"`                         //	Domain configured for this virtual machine.
	HighAvailability            HighAvailability              `json:"high_availability,omitempty"`              //	The virtual machine high availability configuration.
	ID                          string                        `json:"id,omitempty"`                             //	A unique identifier.
	Initialization              Initialization                `json:"initialization,omitempty"`                 //	Reference to the virtual machine’s initialization configuration.
	IO                          Io                            `json:"io,omitempty"`                             //	For performance tuning of IO threading.
	LargeIcon                   Icon                          `json:"large_icon,omitempty"`                     //	Virtual machine’s large icon.
	Lease                       StorageDomainLease            `json:"lease,omitempty"`                          //	Reference to the storage domain this virtual machine/template lease reside on.
	Memory                      int                           `json:"memory,omitempty"`                         //	The virtual machine’s memory, in bytes.
	MemoryPolicy                MemoryPolicy                  `json:"memory_policy,omitempty"`                  //	Reference to virtual machine’s memory management configuration.
	Migration                   MigrationOptions              `json:"migration,omitempty"`                      //	Reference to configuration of migration of running virtual machine to another host.
	MigrationDowntime           int                           `json:"migration_downtime,omitempty"`             //	Maximum time the virtual machine can be non responsive during its live migration to another host in ms.
	MultiQueuesEnabled          bool                          `json:"multi_queues_enabled,omitempty"`           //	If true, each virtual interface will get the optimal number of queues, depending on the available virtual Cpus.
	Name                        string                        `json:"name,omitempty"`                           //	A human-readable name in plain text.
	Origin                      string                        `json:"origin,omitempty"`                         //	The origin of this virtual machine.
	Os                          OperatingSystem               `json:"os,omitempty"`                             //	Operating system type installed on the virtual machine.
	PlacementPolicy             VmPlacementPolicy             `json:"placement_policy,omitempty"`               //	The configuration of the virtual machine’s placement policy.
	RngDevice                   RngDevice                     `json:"rng_device,omitempty"`                     //	Random Number Generator device configuration for this virtual machine.
	SerialNumber                SerialNumber                  `json:"serial_number,omitempty"`                  //	Virtual machine’s serial number in a cluster.
	SmallIcon                   Icon                          `json:"small_icon,omitempty"`                     //	Virtual machine’s small icon.
	SoundcardEnabled            bool                          `json:"soundcard_enabled,omitempty"`              //	If true, the sound card is added to the virtual machine.
	Sso                         Sso                           `json:"sso,omitempty"`                            //	Reference to the Single Sign On configuration this virtual machine is configured for.
	StartPaused                 bool                          `json:"start_paused,omitempty"`                   //	If true, the virtual machine will be initially in 'paused' state after start.
	Stateless                   bool                          `json:"stateless,omitempty"`                      //	If true, the virtual machine is stateless - it’s state (disks) are rolled-back after shutdown.
	Status                      TemplateStatus                `json:"status,omitempty"`                         //	The status of the template.
	StorageErrorResumeBehaviour VmStorageErrorResumeBehaviour `json:"storage_error_resume_behaviour,omitempty"` //	Determines how the virtual machine will be resumed after storage error.
	TimeZone                    TimeZone                      `json:"time_zone,omitempty"`                      //	The virtual machine’s time zone set by oVirt.
	TunnelMigration             bool                          `json:"tunnel_migration,omitempty"`               //	If true, the network data transfer will be encrypted during virtual machine live migration.
	Type                        VmType                        `json:"type,omitempty"`                           //	Determines whether the virtual machine is optimized for desktop or server.
	Usb                         Usb                           `json:"usb,omitempty"`                            //	Configuration of USB devices for this virtual machine (count, type).
	Version                     TemplateVersion               `json:"version,omitempty"`                        //	Indicates whether this is the base version or a sub-version of another template.
	VirtioScsi                  VirtioScsi                    `json:"virtio_scsi,omitempty"`                    //	Reference to VirtIO SCSI configuration.
	VM                          Vm                            `json:"vm,omitempty"`                             //	The virtual machine configuration associated with this template.
}

// TemplateStatus enum
//
// Type representing a status of a virtual machine template.
//   - illegal	-	This status indicates that at least one of the disks of the template is illegal.
//   - locked	-	This status indicates that some operation that prevents other operations with the template is being executed.
//   - ok	-	This status indicates that the template is valid and ready for use.
type TemplateStatus string

// Type representing a version of a virtual machine template.
type TemplateVersion struct {
	VersionName   string `json:"version_name,omitempty"`   //	The name of this version.
	VersionNumber int    `json:"version_number,omitempty"` //	The index of this version in the versions hierarchy of the template.
}
