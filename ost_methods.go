package pncp

import (
	"fmt"
)

func (r *Client) GetListOSTemplates() (Future, error) {
	path := fmt.Sprintf(`/account/%s/node/%s/ostemplate/virtual`, r.AccountID, r.NodeID)
	return r.call(`GET`, path, "", nil)
}

func (r *Client) GetOSTemplateDetails(id uint32) (Future, error) {
	path := fmt.Sprintf(`/ostemplate/%d/virtual`, id)
	return r.call(`GET`, path, "", nil)
}
