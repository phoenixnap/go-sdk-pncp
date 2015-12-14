package pncp

import (
	"fmt"
)

func (r *Client) ListVirtualMachinesByAccount() (Future, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine`, r.AccountID)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) ListVirtualMachinesByNode() (Future, error) {
	path := fmt.Sprintf(`/account/%s/node/%s/device/virtualmachine`, r.AccountID, r.NodeID)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetVirtualMachineDetails(id uint64) (Future, error) {
	path := fmt.Sprintf(`/virtualmachine/%s`, id)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetVirtualMachineResourceDetails(resource string) (Future, error) {
	return r.call(`GET`, resource, "", nil)
}

func (r *Client) CreateVirtualMachine(props CreateVMRequest) (Future, error) {
	path := fmt.Sprintf(`/account/%s/node/%s/device/virtualmachine`, r.AccountID, r.NodeID)
	return r.call(`POST`, path, "", props)
}

func (r *Client) SetVirtualMachinePowerState(id uint64, state string) (Future, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/power`, state)
	return r.SetVirtualMachineResourcePowerState(path, state)
}

func (r *Client) SetVirtualMachineResourcePowerState(resource, state string) (Future, error) {
	if state != `on` && state != `off` {
		panic(`Invalid power state provided`)
	}
	qs := fmt.Sprintf("?powerState=%s", state)
	return r.call(`PUT`, resource, qs, nil)
}

func (r *Client) RebootVirtualMachine(id uint64) (Future, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/reboot`, id)
	return r.call(`PUT`, path, "", nil)
}

func (r *Client) RebootVirtualMachineResource(resource string) (Future, error) {
	path := fmt.Sprintf(`%s/reboot`, resource)
	return r.call(`PUT`, path, "", nil)
}

func (r *Client) CloneVirtualMachine(id uint64) (Future, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/clone`, id)
	return r.call(`POST`, path, "", nil)
}

func (r *Client) ModifyVirtualMachine(id uint64, props ModifyVMRequest) (Future, error) {
	path := fmt.Sprintf(`/virtualmachine/%s`, id)
	return r.call(`POST`, path, "", props)
}

func (r *Client) DeleteVirtualMachine(id uint64, releaseIP bool) (Future, error) {
	path := fmt.Sprintf(`/virtualmachine/{vmID}`, id)
	return r.DeleteVirtualMachineResource(path, releaseIP)
}

func (r *Client) DeleteVirtualMachineResource(resource string, releaseIP bool) (Future, error) {
	qs := fmt.Sprintf("?releaseIPs=%s", releaseIP)
	return r.call(`PUT`, resource, qs, nil)
}

func (r *Client) GetVirtualMachineTags(id uint64) (Future, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine/%s/tags`, r.AccountID, id)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) AddTagToVirtualMachine(id uint64, tag string) (Future, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine/%s/tags/%s`, r.AccountID, id, tag)
	return r.call(`PUT`, path, "", nil)
}

func (r *Client) RemoveTagFromVirtualMachine(id uint64, tag string) (Future, error) {
	path := fmt.Sprintf(`/account/%s/virtualmachine/%s/tags/%s`, r.AccountID, id, tag)
	return r.call(`DELETE`, path, "", nil)
}

func (r *Client) Version() string {
	return libVersion
}
