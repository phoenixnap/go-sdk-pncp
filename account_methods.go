package pncp

import (
	"fmt"
)

func (r *Client) GetAccountDetails() (Future, string, bool, uint64, error) {
	path := fmt.Sprintf(`/account/%s`, r.AccountID)
	return r.call(`GET`, path, "", nil)
}
