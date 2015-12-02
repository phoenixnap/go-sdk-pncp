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
	Result                interface{}
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
	response    *SyncResponse
	api         *Client
}

func (ar AsyncResponse) Get(r interface{}) (err error) {
	if ar.api == nil {
		return errors.New(`Client API is unset.`)
	} else if ar.ResourceURL == "" {
		return errors.New(`The resource to poll is unset.`)
	}
	for {
		if ar.response != nil {
			return ar.response.Get(r)
		}

		var (
			out       Future
			emsg      string
			retriable bool
		)

		result := &Task{}
		out, emsg, retriable, _, err = ar.api.call(`GET`, ar.ResourceURL, ``, ``)
		if err != nil {
			return
		}
		if emsg != `` && !retriable {
			err = errors.New(emsg)
			return
		}

		out.Get(result) // Unmarshal the task response (we know its a synchronous call)
		if err != nil {
			return err
		}
		if result.RequestStateEnum == `CLOSED_SUCCESSFUL` || result.RequestStateEnum == `CLOSED_FAILED` {
			r = result.Result

			return
		}
		time.Sleep(ar.api.Backoff)
	}
}
func (ar AsyncResponse) TimedGet(r interface{}, ttl time.Duration) (err error) {
	// TODO: Implement timeout
	return nil
}
