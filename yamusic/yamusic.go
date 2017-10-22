package yamusic

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiURL = "https://api.music.yandex.net"
)

type (
	// Doer is an interface that can do http request
	Doer interface {
		Do(req *http.Request) (*http.Response, error)
	}
	// A Client manages communication with the Yandex.Music API.
	Client struct {
		// HTTP client used to communicate with the API.
		client Doer
		// Base URL for API requests.
		baseURL *url.URL
		// Access token to Yandex.Music API
		accessToken string
		userID      int
		// Services
		genres    *GenresService
		search    *SearchService
		account   *AccountService
		feed      *FeedService
		playlists *PlaylistsService
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
	c.account = &AccountService{client: c}
	c.feed = &FeedService{client: c}
	c.playlists = &PlaylistsService{client: c}

	return c
}

// HTTPClient sets http client for Yandex.Music client
func HTTPClient(httpClient Doer) func(*Client) {
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

// AccessToken sets user_id and access token for Yandex.Music client
func AccessToken(userID int, accessToken string) func(*Client) {
	return func(c *Client) {
		if userID != 0 {
			c.userID = userID
		}

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
func (c *Client) NewRequest(
	method,
	urlStr string,
	body interface{},
) (*http.Request, error) {

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	var reader io.Reader
	if body != nil {
		switch v := body.(type) {
		case url.Values:
			reader = strings.NewReader(v.Encode())
		default:
			buf := new(bytes.Buffer)
			err = json.NewEncoder(buf).Encode(body)
			if err != nil {
				return nil, err
			}

			reader = buf
		}
	}

	req, err := http.NewRequest(method, u.String(), reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "OAuth "+c.accessToken)
	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(
	ctx context.Context,
	req *http.Request,
	v interface{},
) (*http.Response, error) {

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := resp.Body.Close()
		if closeErr != nil {
			log.Println("close response body error: ", closeErr)
		}
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
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

// Account returns account service
func (c *Client) Account() *AccountService {
	return c.account
}

// Feed returns feed service
func (c *Client) Feed() *FeedService {
	return c.feed
}

// Playlists returns feed service
func (c *Client) Playlists() *PlaylistsService {
	return c.playlists
}
