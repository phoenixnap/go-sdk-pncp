package pncp

import (
	"fmt"
)

func (r *Client) ListVirtualMachinesByAccount() (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine`, r.AccountID)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) ListVirtualMachinesByNode() (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s/node/%s/device/virtualmachine`, r.AccountID, r.NodeID)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetVirtualMachineDetails(id uint64) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s`, id)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetVirtualMachineDetailsURL(url string) (Future, string, bool, uint64, error) {
	return r.call(`GET`, url, "", nil)
}

func (r *Client) CreateVirtualMachine(props CreateVMRequest) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s/node/%s/device/virtualmachine`, r.AccountID, r.NodeID)
	return r.call(`POST`, path, "", props)
}

func (r *Client) SetVirtualMachinePowerState(state string) (Future, string, bool, uint64, error) {
	if state != `on` && state != `off` {
		panic(`Invalid power state provided`)
	}
	path := fmt.Sprintf(`/virtualmachine/%s/power`, state)
	qs := fmt.Sprintf("?powerState=%s", state)
	return r.call(`PUT`, path, qs, nil)
}

func (r *Client) RebootVirtualMachine(id uint64) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/reboot`, id)
	return r.call(`PUT`, path, "", nil)
}

func (r *Client) CloneVirtualMachine(id uint64) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/clone`, id)
	return r.call(`POST`, path, "", nil)
}

func (r *Client) ModifyVirtualMachine(id uint64, props ModifyVMRequest) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s`, id)
	return r.call(`POST`, path, "", props)
}

func (r *Client) DeleteVirtualMachine(id uint64, releaseIP bool) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/{vmID}`, id)
	qs := fmt.Sprintf("?releaseIPs=%s", releaseIP)
	return r.call(`PUT`, path, qs, nil)
}

func (r *Client) GetVirtualMachineTags(id uint64) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine/%s/tags`, r.AccountID, id)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) AddTagToVirtualMachine(id uint64, tag string) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine/%s/tags/%s`, r.AccountID, id, tag)
	return r.call(`PUT`, path, "", nil)
}

func (r *Client) RemoveTagFromVirtualMachine(id uint64, tag string) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine/%s/tags/%s`, r.AccountID, id, tag)
	return r.call(`DELETE`, path, "", nil)
}

func (r *Client) Version() string {
	return libVersion
}
