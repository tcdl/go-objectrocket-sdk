package objectrocket

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	// Default URL for ObjectRocket API
	defaultBaseURL = "https://sjc-api.objectrocket.com/v2"
)

// A APIClient manages communication with the ObjectRocketAPI.
type Client struct {
	// client is the HTTP client used to communicate with the API.
	client *http.Client
	// BaseURL for API requests.
	BaseURL *url.URL
	// APIKey for API requests
	APIKey string
	// Sevices used for communications with API
	Instance InstanceService
}

// NewAPIClient returns a new ObjectRocket API client. If a nil httpClient
// is provided, http.DefaultClient will be used. To use API methods which
// require authentication, provide an http.Client that will perform the
// authentication for you .
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, APIKey: apiKey}

	c.Inst = InstService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the APIClient.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Auth-Token", c.APIKey)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := response.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err := io.Copy(w, response.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err := json.NewDecoder(response.Body).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return response, err
}
