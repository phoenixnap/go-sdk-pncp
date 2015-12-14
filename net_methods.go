package pncp

import (
	"fmt"
)

func (r *Client) GetNetworkConfiguration() (Future, error) {
	path := fmt.Sprintf(`/account/%s/node/%s/network`, r.AccountID, r.NodeID)
	return r.call(`GET`, path, "", nil)
}