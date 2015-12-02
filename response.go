package pncp

import (
	"encoding/json"
	"errors"
	"time"
)

type Future interface {
	Get(r interface{}) (err error)
	TimedGet(r interface{}, ttl time.Duration) (err error)
}

type Task struct {
	PercentageComplete    int
	RequestStateEnum      string
	ProcessDescription    string
	LatestTaskDescription string
	Result                Resource
	ErrorCode             string
	ErrorMessage          string
	LastUpdatedTimestamp  string
	CreatedTimestamp      string
}

//\\//\\//\\//\\//\\//\\//\\//\\//\\
// Synchrounous Implementation
//\\//\\//\\//\\//\\//\\//\\//\\//\\

type SyncResponse struct {
	body []byte
}

func (sr SyncResponse) Get(r interface{}) error {
	// Unmarshal into given type now
	return json.Unmarshal(sr.body, r)
}
func (sr SyncResponse) TimedGet(r interface{}, ttl time.Duration) (err error) {
	return sr.Get(r)
}

//\\//\\//\\//\\//\\//\\//\\//\\//\\
// Asynchronous Implementation
//\\//\\//\\//\\//\\//\\//\\//\\//\\

type AsyncResponse struct {
	ResourceURL string `json:"resourceURL"`
	response    *Resource
	api         *Client
}

func (ar AsyncResponse) Get(r interface{}) error {
	var (
		rr *Resource
		ok bool
	)
	if rr, ok = r.(*Resource); !ok {
		return errors.New(`This function only binds to Resources.`)
	}

	if ar.api == nil {
		return errors.New(`Client API is unset.`)
	} else if ar.ResourceURL == "" {
		return errors.New(`The resource to poll is unset.`)
	} else if ar.response != nil {
		rr.URL = ar.response.URL
	}

	// The response has not been retrieved and conditions are correct for retrieval
	for {
		// Poll for the task status
		out, emsg, retriable, _, err := ar.api.call(`GET`, ar.ResourceURL, ``, ``)
		if err != nil {
			return err
		}
		if emsg != `` && !retriable {
			return errors.New(emsg)
		}

		// Unmarshall the task
		resp := &Task{}
		out.Get(resp) // Unmarshal the task response (we know its a synchronous call)
		if err != nil {
			return err
		}
		if resp.RequestStateEnum == `CLOSED_SUCCESSFUL` || resp.RequestStateEnum == `CLOSED_FAILED` {
			ar.response = &resp.Result
			rr.URL = ar.response.URL
			return nil
		}
		time.Sleep(ar.api.Backoff)
	}
}
func (ar AsyncResponse) TimedGet(r interface{}, ttl time.Duration) (err error) {
	// TODO: Implement timeout
	return nil
}
