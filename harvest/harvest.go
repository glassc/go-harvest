package harvest

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
	"net/url"
	"io"
)

const (
 	defaultBaseUrl = "https://%v.harvestapp.com" 
	mediaType = "application/json"
)

type service struct {
	client *Client
}

type Client struct {
	client   *http.Client
	baseURL  *url.URL

	Account *AccountService
}

func NewClient(account string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(fmt.Sprintf(defaultBaseUrl, account))
	c := &Client{client: httpClient, baseURL: baseUrl}
	c.Account = &AccountService{c}
	return c
}

type BasicAuthTransport struct {
	Username string
	Password string
}


func (t *BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {	
	req.SetBasicAuth(t.Username, t.Password)
	return http.DefaultTransport.RoundTrip(req)
}

// Client returns an *http.Client that makes requests that are authenticated
// using HTTP Basic Authentication.
func (t *BasicAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}


func (c *Client) newRequest(method string, action string, body interface{}) (*http.Request, error) {
	actionUrl, _ := url.Parse(action)
	url := c.baseURL.ResolveReference(actionUrl)

	var req *http.Request
	var err error

	var encodedBody io.ReadWriter
	if body != nil {
		encodedBody = new(bytes.Buffer)
		err := json.NewEncoder(encodedBody).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err = http.NewRequest(method, url.String(), encodedBody)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", mediaType)
	req.Header.Set("Accept", mediaType)
	return req, nil
}

func (c *Client) get(action string) (*http.Response, error) {
	req, err := c.newRequest("GET", action, nil)

	if err != nil {
		return nil, err
	}


	return c.client.Do(req)
}