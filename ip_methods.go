package pncp

import (
	"fmt"
)

func (r *Client) ListPublicIPsForVirtualMachine(id uint64) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/publicip`, id)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) ListPrivateIPsForVirtualMachine(id uint64) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/privateip`, id)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetPublicIPDetailsOnVirtualMachine(id uint64, ip string) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/publicip/%s`, id, ip)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetPrivateIPDetailsOnVirtualMachine(id uint64, ip string) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/privateip/%s`, id, ip)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) AssignPublicIPToVirtualMachine(id uint64, spec PublicIPSpec) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/publicip`, id)
	return r.call(`POST`, path, "", spec)
}

func (r *Client) AssignPrivateIPToVirtualMachine(id uint64, spec PrivateIPSpec) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/privateip`, id)
	return r.call(`POST`, path, "", spec)
}

func (r *Client) ModifyPublicIPOnVirtualMachine(id uint64, ip string, spec PublicIPUpdateSpec) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/publicip/%s`, id, ip)
	return r.call(`PUT`, path, "", spec)
}

func (r *Client) ModifyPrivateIPOnVirtualMachine(id uint64, ip string, spec PrivateIPUpdateSpec) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/privateip/%s`, id, ip)
	return r.call(`PUT`, path, "", spec)
}

func (r *Client) ReleasePublicIPOnVirtualMachine(id uint64, ip string, release bool) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/publicip/%s`, id, ip)
	qs := fmt.Sprintf(`?moveToReservedList=%s`, release)
	return r.call(`DELETE`, path, qs, nil)
}

func (r *Client) ReleasePrivateIPOnVirtualMachine(id uint64, ip string) (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/virtualmachine/%s/privateip/%s`, id, ip)
	return r.call(`DELETE`, path, "", nil)
}
