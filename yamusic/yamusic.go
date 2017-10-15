package yamusic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	apiURL = "https://api.music.yandex.net"
)

type (
	// A Client manages communication with the Yandex.Music API.
	Client struct {
		// HTTP client used to communicate with the API.
		client *http.Client
		// Base URL for API requests.
		baseURL *url.URL
		// Access token to Yandex.Music API
		accessToken string
		// Services
		genres *GenresService
		search *SearchService
	}
)

// NewClient returns a new API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
func NewClient(options ...func(*Client)) *Client {
	baseURL, _ := url.Parse(apiURL)

	c := &Client{
		client:  http.DefaultClient,
		baseURL: baseURL,
	}

	for _, option := range options {
		option(c)
	}

	c.genres = &GenresService{client: c}
	c.search = &SearchService{client: c}

	return c
}

// HTTPClient sets http client for Yandex.Music client
func HTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		if httpClient != nil {
			c.client = httpClient
		}
	}
}

// BaseURL sets base API URL for Yandex.Music client
func BaseURL(baseURL *url.URL) func(*Client) {
	return func(c *Client) {
		if baseURL != nil {
			c.baseURL = baseURL
		}
	}
}

// AccessToken sets access token for Yandex.Music client
func AccessToken(accessToken string) func(*Client) {
	return func(c *Client) {
		if accessToken != "" {
			c.accessToken = accessToken
		}
	}
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}

// Genres returns genres service
func (c *Client) Genres() *GenresService {
	return c.genres
}

// Search returns genres service
func (c *Client) Search() *SearchService {
	return c.search
}
