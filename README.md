# ovirt_inventory

App for getting statistic from oVirt cluster

## Requirements

- oVirt 4.3

## General

App fetch VMs stats with corresponding disks size

### Stats:
- Comment       - Free text containing comments about this object.
- Cpu           - Numder of CPU for the virtual machine.
- CreationTime  - The virtual machine creation date. When requesting the JSON representation the engine uses a different, format:- an integer containing the number of seconds since Jan 1st 1970, also know as epoch time [here](https://en.wikipedia.org/wiki/Unix_time).
- Description   - A human-readable description in plain text.
- FQDN          - Fully qualified domain name of the virtual machine.
- ID            - A unique identifier.
- Memory        - The virtual machine’s memory, in bytes.
- Name          - A human-readable name in plain text.
- OS            - Operating system type installed on the virtual machine. From the field Type from struct OperatingSystem
- RunOnce       - If true, the virtual machine has been started using the run once command, meaning it’s configuration might differ from the stored one for the purpose of this single run.
- SerialNumber  - Virtual machine’s serial number in a cluster. The field Value of struct SerialNumber
- StartTime     - The date in which the virtual machine was started.
- Status        - The current status of the virtual machine.
- StatusDetail  - Human readable detail of current status.
- StopReason    - The reason the virtual machine was stopped.
- StopTime      - The date in which the virtual machine was stopped.
- HddDiskSize   - The size of all disks attached to vm. Sum of initial_size from Disk struct, in bytes. Group by StorageType (HDD or SSD).
- SsdDiskSize   - The size of all disks attached to vm. Sum of initial_size from Disk struct, in bytes. Group by StorageType (HDD or SSD).
- VmDisksCount  - The count of attached virtual disks
