package pncp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	//"crypto/tls"
)

// Generic call
func (r *Client) call(method, path, qs string, inBody interface{}) (Future, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response

		authContext AuthContext
		reqBody     []byte
		url         string
		err         error
	)

	// Preconditions and sanitize input

	// Construct the HTTP Request
	if inBody != nil {
		reqBody, err = json.Marshal(inBody)
		if err != nil {
			return nil, err
		}
	} else {
		reqBody = nil
	}

	// Construct URL and AuthContext

	url = fmt.Sprintf("%s%s%s", r.Endpoint, path, qs)
	authContext = NewAuthContext(method, path, qs, r.ApplicationKey, r.SharedSecret)

	if r.Debug {
		log.Printf("CALL\n\tURL: %s\n\tMethod: %s\n\tRequestBody? %t\n", url, method, inBody != nil)
		log.Printf("AUTH\n")
		log.Printf("\tMethod: %s\n", authContext.Method)
		log.Printf("\tResourcePath: %s\n", authContext.ResourcePath)
		log.Printf("\tQueryString: %s\n", authContext.QueryString)
		log.Printf("\tApplicationKey: %s\n", authContext.ApplicationKey)
		log.Printf("\tStringToSign: %s\n", authContext.StringToSign)
		log.Printf("\tRequestSignature: %s\n", authContext.RequestSignature)
		log.Printf("\tEncodedCredentials: %s\n", authContext.EncodedCredentials)
		log.Printf("\tAuthenticator: %s\n", authContext.Authenticator)
		if inBody != nil {
			log.Printf("BODY:\n\t%s", reqBody)
		}
	}

	switch method {
	case `GET`:
		fallthrough
	case `PUT`:
		fallthrough
	case `POST`:
		fallthrough
	case `DELETE`:
		req, err = constructRequest(method, url, reqBody)
	default:
		err = errors.New(`Unknown method`)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", documentType)
	req.Header.Set("Content-Type", documentType)
	req.Header.Set("Authorization", authContext.Authenticator)

	// temp ignore bad TLS
	tr := &http.Transport{
	//TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if r.Debug {
		log.Printf("Response status: %s code: %d", resp.Status, resp.StatusCode)
	}

	if resp.StatusCode == 200 {
		rawOut, e := ioutil.ReadAll(resp.Body)
		if r.Debug {
			log.Printf("Response body: %s", rawOut)
		}
		if e != nil {
			return nil, &APIError{error: e, Retriable: true}
		}
		return &SyncResponse{body: rawOut}, nil
	} else if resp.StatusCode == 202 {
		rawOut, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			return nil, &APIError{error: e, Retriable: true}
		}

		res := &Resource{}
		e = json.Unmarshal(rawOut, res)
		if e != nil {
			return nil, &APIError{error: e, Retriable: true}
		}
		return &AsyncResponse{api: r, ResourceURL: res.URL}, nil
	}

	if resp.StatusCode == 500 {
		eref, _ := strconv.ParseUint(resp.Header.Get("X-Application-Error-Reference"), 10, 64)
		return nil, &APIError{
			error:     errors.New(resp.Header.Get("X-Application-Error-Description")),
			Retriable: true,
			Eref:      eref,
		}
	} else if resp.StatusCode == 400 {
		eref, _ := strconv.ParseUint(resp.Header.Get("X-Application-Error-Reference"), 10, 64)
		return nil, &APIError{
			error:     errors.New(resp.Header.Get("X-Application-Error-Description")),
			Retriable: false,
			Eref:      eref,
		}
	} else if resp.StatusCode == 401 {
		eref, _ := strconv.ParseUint(resp.Header.Get("X-Application-Error-Reference"), 10, 64)
		return nil, &APIError{
			error:     errors.New(resp.Header.Get("X-Application-Error-Description")),
			Retriable: false,
			Eref:      eref,
		}
	} else {
		return nil, &APIError{
			error:     errors.New(resp.Status),
			Retriable: false,
		}
	}

}

func constructRequest(method, url string, body []byte) (*http.Request, error) {
	if body == nil {
		return http.NewRequest(method, url, nil)
	} else {
		return http.NewRequest(method, url, bytes.NewBuffer(body))
	}
}
