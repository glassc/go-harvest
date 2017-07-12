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

type Client struct {
	client   *http.Client
	baseURL  *url.URL
}

func NewClient(httpClient *http.Client, account string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseUrl, _ := url.Parse(fmt.Sprintf(defaultBaseUrl, account))
	c := &Client{client: httpClient, baseURL: baseUrl}
  
	return c
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

func (c *Client) get(action string, body interface{}) (*http.Response, error) {
	req, err := c.newRequest("GET", action, body)

	if err != nil {
		return nil, err
	}


	return c.client.Do(req)
}