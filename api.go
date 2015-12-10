package pncp

import (
	"time"
)

const (
	libVersion   = "0.1a"
	documentType = "application/vnd.pncp.v.1.0+json"
)

// Interfaces

type API interface {

	// Account Resources
	GetAccountDetails() (Future, error)

	// Virtual Machine Resources
	ListVirtualMachinesByAccount() (Future, error)
	ListVirtualMachinesByNode() (Future, error)
	GetVirtualMachineDetails(id uint64) (Future, error)
	GetVirtualMachineResourceDetails(resource string) (Future, error)
	CreateVirtualMachine(props CreateVMRequest) (Future, error)
	SetVirtualMachinePowerState(state string) (Future, error)
	RebootVirtualMachine(id uint64) (Future, error)
	RebootVirtualMachineResource(resource string) (Future, error)
	CloneVirtualMachine(id uint64) (Future, error)
	ModifyVirtualMachine(id uint64, props ModifyVMRequest) (Future, error)
	DeleteVirtualMachine(id uint64, releaseIP bool) (Future, error)
	DeleteVirtualMachineResource(resource string, releaseIP bool) (Future, error)
	GetVirtualMachineTags(id uint64) (Future, error)
	AddTagToVirtualMachine(id uint64, tag string) (Future, error)
	RemoveTagFromVirtualMachine(id uint64, tag string) (Future, error)

	// Network Configuration Services
	GetNetworkConfiguration() (Future, error)

	// IP Management Resources
	ListPublicIPsForVirtualMachine(id uint64) (Future, error)
	ListPublicIPsForVirtualMachineResource(resource string) (Future, error)
	ListPrivateIPsForVirtualMachine(id uint64) (Future, error)
	GetPublicIPDetailsOnVirtualMachine(id uint64, ip string) (Future, error)
	GetPublicIPResourceDetails(resource string) (Future, error)
	GetPrivateIPDetailsOnVirtualMachine(id uint64, ip string) (Future, error)
	AssignPublicIPToVirtualMachine(id uint64, spec PublicIPSpec) (Future, error)
	AssignPublicIPToVirtualMachineResource(resource string, spec PublicIPSpec) (Future, error)
	AssignPrivateIPToVirtualMachine(id uint64, spec PrivateIPSpec) (Future, error)
	ModifyPublicIPOnVirtualMachine(id uint64, ip string, spec PublicIPUpdateSpec) (Future, error)
	ModifyPrivateIPOnVirtualMachine(id uint64, ip string, spec PrivateIPUpdateSpec) (Future, error)
	ReleasePublicIPOnVirtualMachine(id uint64, ip string, release bool) (Future, error)
	ReleasePrivateIPOnVirtualMachine(id uint64, ip string) (Future, error)

	// Bare Metal Resources
	/*
		ListBareMetalDevicesByAccount() (Future, string, bool, uint64, error)
		ListBareMetalDevicesByNode() (Future, string, bool, uint64, error)
		GetBareMetalDeviceTraits(id uint64) (Future, string, bool, uint64, error)
		ModifyBareMetalDevice(id uint64, name, description string) (Future, string, bool, uint64, error)
		SetBareMetalDevicePowerState(id uint64, state string) (Future, string, bool, uint64, error)
		RebootBareMetalDevice(id uint64) (Future, string, bool, uint64, error)
		GetBareMetalDeviceAssignment(id uint64) (Future, string, bool, uint64, error)
		GetBareMetalDeviceAssignmentDetails(id uint64) (Future, string, bool, uint64, error)
		AssignBareMetalDevice(id uint64) (Future, string, bool, uint64, error)
		RemoveBareMetalDeviceAssignment(id uint64) (Future, string, bool, uint64, error)
		GetBareMetalDeviceTags(id uint64) (Future, string, bool, uint64, error)
		AddTagToBareMetalDevice(id uint64, tag string) (Future, string, bool, uint64, error)
		RemoveTagFromBareMetalDevice(id uint64, tag string) (Future, string, bool, uint64, error)
	*/
	// Firewall Resources

	// OS Template Resources
	GetListOSTemplates() (Future, error)
	GetOSTemplateDetails(id uint32) (Future, error)

	Version() string
}

const (
	PowerOn                      = `on`
	PowerOff                     = `off`
	BillingMethodUbersmith       = `UBERSMITH`
	BillingMethodExternal        = `EXTERNAL`
	AccountStatusGoodStanding    = `GOOD_STANDING`
	AccountStatusUsersSuspended  = `USERS_SUSPENDED`
	AccountStatusOnHold          = `ON_HOLD`
	AccountStatusCancelRequested = `CANCEL_REQUESTED`
	AccountStatusCancelled       = `CANCELLED`
	AccountStatusTerminated      = `TERMINATED`
	AccountStatusEndClient       = `END_CLIENT`
)

// API implementation

type Client struct {
	Endpoint       string
	AccountID      string
	ApplicationKey string
	SharedSecret   string
	NodeID         string
	Debug          bool
	Backoff        time.Duration
}

func NewClient(endpoint, accountid, key, secret, node string, debug bool) *Client {
	return &Client{
		Endpoint:       endpoint,
		AccountID:      accountid,
		ApplicationKey: key,
		SharedSecret:   secret,
		NodeID:         node,
		Debug:          debug,
		Backoff:        time.Duration(3) * time.Second,
	}
}

type APIError struct {
	error
	Retriable bool
	Eref      uint64
}

//
// API Request/Response Structures
//

type AccountDetails struct {
	Name                   string             `json:"name"`
	Email                  string             `json:"email"`
	Description            string             `json:"description"`
	AdminURL               string             `json:"adminUrl"`
	ClientAssignedID       string             `json:"clientAssignedId"`
	ReportBugEmail         string             `json:"reportBugEmail"`
	AccountStatus          string             `json:"accountStatus"`
	Account                Resource           `json:"accountResource"`
	SignUpDate             string             `json:"signUpDate"`
	ParentAccount          Resource           `json:"parentAccountResouce"`
	PrimaryContact         Resource           `json:"primaryContactResource"`
	TechnicalContact       Resource           `json:"technicalContactResource"`
	BillingMethod          string             `json:"billingMethod"`
	AssignedPricingProfile Resource           `json:"assignedPricingProfileResource"`
	Permissions            AccountPermissions `json:"permissions"`
}
type AccountPermissions struct {
	Reseller  bool `json:"reseller"`
	Virtual   bool `json:"virtual"`
	BareMetal bool `json:"bareMetal"`
}

type Resource struct {
	URL string `json:"resourceURL"`
}

type ResourceList []Resource

type CreateVMRequest struct {
	Name                    string   `json:"name"`
	Description             string   `json:"description"`
	StorageInGB             uint16   `json:"storageGB"`
	MemoryInMB              uint32   `json:"memoryMB"`
	VCpuCount               uint8    `json:"vCPUs"`
	StorageType             string   `json:"storageType,omitempty"`
	PowerStatus             string   `json:"powerStatus,omitempty"`
	OperatingSystemTemplate Resource `json:"operatingSystemTemplate,omitempty"`
	ImageResource           string   `json:"imageResource,omitempty"`
	Password                string   `json:"newOperatingSystemAdminPassword,omitempty"`
}

type ModifyVMRequest struct {
	Description string `json:"description"`
	MemoryInMB  uint32 `json:"memoryMB"`
	VCpuCount   uint8  `json:"vCPUs"`
}

type VirtualMachineDetails struct {
	Name                    string     `json:"name"`
	Description             string     `json:"description"`
	StorageInGB             uint16     `json:"storageGB"`
	MemoryInMB              uint32     `json:"memoryMB"`
	VCpuCount               uint8      `json:"vCPUs"`
	StorageType             string     `json:"storageType,omitempty"`
	PowerStatus             string     `json:"powerStatus,omitempty"`
	OperatingSystemTemplate Resource   `json:"operatingSystemTemplate,omitempty"`
	ImageResource           string     `json:"imageResource,omitempty"`
	NodeResource            Resource   `json:"nodeResource,omitempty"`
	AccountResource         Resource   `json:"accountResource,omitempty"`
	Disks                   []Resource `json:"disks,omitempty"`
	MACAddress              string     `json:"macAddress,omitempty"`
	DeducedPrivateIPs       []string   `json:"deducedPrivateIps,omitempty"`
	IPMappings              []VMIPMap  `json:"ipMappings,omitempty"`
}

type OSTemplate struct {
	Name                         string `json:"name"`
	Version                      string `json:"version"`
	MinimumStorageSpaceInGB      uint16 `json:"minimumStorageSpace"`
	DefaultAdministratorUsername string `json:"defaultAdministratorUsername"`
	DiskExpandable               bool   `json:"diskExpandable"`
}

type VMIPMap struct {
	PrivateIP string
	PublicIPs []string
}

type PublicIPSpec struct {
	IPFromReserved   string `json:"ipFromReserved,omitempty"`
	PrivateIPMapping string `json:"privateIpMapping,omitempty"`
}

type PublicIPUpdateSpec struct {
	PrivateIPMapping string `json:"privateIpMapping"`
}

type PrivateIPSpec struct {
	IPAddress       string   `json:"ipAddress"`
	PublicIPMapping []string `json:"publicIpMappings"`
}

type PrivateIPUpdateSpec struct {
	PublicIPMapping []string `json:"publicIpMappings"`
}

type PublicIPAssignmentDesc struct {
	IPAddress    string     `json:"ipAddress,omitempty"`
	Type         string     `json:"ipType,omitempty"`
	Reserved     bool       `json:"reserved,omitempty"`
	AssignedToVM []Resource `json:"assignedTo,omitempty"`
	PrivateIP    string     `json:"privateIPMapping,omitempty"`
	Node         Resource   `json:"nodeResource,omitempty"`
	Account      Resource   `json:"accountResouce,omitempty"`
}

type PrivateIPAssignmentDesc struct {
	IPAddress     string   `json:"ipAddress,omitempty"`
	AssignedToVM  Resource `json:"assignedTo,omitempty"`
	PublicMapping []string `json:"publicIPMapping,omitempty"`
	Node          Resource `json:"nodeResource,omitempty"`
	Account       Resource `json:"accountResouce,omitempty"`
}

type NetworkConfigurationDesc struct {
	Mode                    string   `json:"networkConfig,omitempty"`
	NetworkStoragePrivateIP string   `json:"networkStoragePrivateIP,omitempty"`
	Mask                    string   `json:"networkMask,omitempty"`
	GatewayIP               string   `json:"gatewayIP,omitempty"`
	DNSServers              []string `json:"dnsServers,omitempty"`
	PrivateIPRangeFirst     string   `json:"privateIpRangeFirst,omitempty"`
	PrivateIPRangeLast      string   `json:"privateIpRangeLast,omitempty"`
	LeaseTimeDefaultSec     int64    `json:"leaseTimeDefaultSecs,omitempty"`
	LeaseTimeMaxSecs        int64    `json:"leaseTimeMaxSecs,omitempty"`
}
