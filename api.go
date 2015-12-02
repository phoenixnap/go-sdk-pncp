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
	GetAccountDetails() (Future, string, bool, uint64, error)

	// Virtual Machine Resources
	ListVirtualMachinesByAccount() (Future, string, bool, uint64, error)
	ListVirtualMachinesByNode() (Future, string, bool, uint64, error)
	GetVirtualMachineDetails(id uint64) (Future, string, bool, uint64, error)
	GetVirtualMachineDetailsURL(url string) (Future, string, bool, uint64, error)
	CreateVirtualMachine(props CreateVMRequest) (Future, string, bool, uint64, error)
	SetVirtualMachinePowerState(state string) (Future, string, bool, uint64, error)
	RebootVirtualMachine(id uint64) (Future, string, bool, uint64, error)
	CloneVirtualMachine(id uint64) (Future, string, bool, uint64, error)
	ModifyVirtualMachine(id uint64, props ModifyVMRequest) (Future, string, bool, uint64, error)
	DeleteVirtualMachine(id uint64, releaseIP bool) (Future, string, bool, uint64, error)
	GetVirtualMachineTags(id uint64) (Future, string, bool, uint64, error)
	AddTagToVirtualMachine(id uint64, tag string) (Future, string, bool, uint64, error)
	RemoveTagFromVirtualMachine(id uint64, tag string) (Future, string, bool, uint64, error)

	// IP Management Resources
	ListPublicIPsForVirtualMachine(id uint64) (Future, string, bool, uint64, error)
	ListPrivateIPsForVirtualMachine(id uint64) (Future, string, bool, uint64, error)
	GetPublicIPDetailsOnVirtualMachine(id uint64, ip string) (Future, string, bool, uint64, error)
	GetPrivateIPDetailsOnVirtualMachine(id uint64, ip string) (Future, string, bool, uint64, error)
	AssignPublicIPToVirtualMachine(id uint64, spec PublicIPSpec) (Future, string, bool, uint64, error)
	AssignPrivateIPToVirtualMachine(id uint64, spec PrivateIPSpec) (Future, string, bool, uint64, error)
	ModifyPublicIPOnVirtualMachine(id uint64, ip string, spec PublicIPUpdateSpec) (Future, string, bool, uint64, error)
	ModifyPrivateIPOnVirtualMachine(id uint64, ip string, spec PrivateIPUpdateSpec) (Future, string, bool, uint64, error)
	ReleasePublicIPOnVirtualMachine(id uint64, ip string, release bool) (Future, string, bool, uint64, error)
	ReleasePrivateIPOnVirtualMachine(id uint64, ip string) (Future, string, bool, uint64, error)

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
	GetListOSTemplates() (Future, string, bool, uint64, error)
	GetOSTemplateDetails(id uint32) (Future, string, bool, uint64, error)

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
		Backoff:        time.Duration(10) * time.Second,
	}
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
	IPFromReserved   string `json:"ipFromReserved"`
	PrivateIPMapping string `json:"privateIpMapping"`
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
	IPAddress    string   `json:"ipAddress"`
	Type         string   `json:"ipType"`
	Reserved     bool     `json:"reserved"`
	AssignedToVM Resource `json:"assignedTo"`
	PrivateIP    string   `json:"privateIPMapping"`
	Node         Resource `json:"nodeResource"`
	Account      Resource `json:"accountResouce"`
}

type PrivateIPAssignmentDesc struct {
	IPAddress     string   `json:"ipAddress"`
	AssignedToVM  Resource `json:"assignedTo"`
	PublicMapping []string `json:"publicIPMapping"`
	Node          Resource `json:"nodeResource"`
	Account       Resource `json:"accountResouce"`
}
