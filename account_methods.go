package pncp

import (
	"fmt"
)

func (r *Client) GetAccountDetails() (Future, error) {
	path := fmt.Sprintf(`/account/%s`, r.AccountID)
	return r.call(`GET`, path, "", nil)
}
