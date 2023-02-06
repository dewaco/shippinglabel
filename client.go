package shippinglabel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dewaco/shippinglabel/types"
	"golang.org/x/net/http2"
	"io"
	"net/http"
	"net/url"
	"time"
)

const productionURL = "https://api.shippinglabel.de/v2"
const developmentURL = "https://api.dev.shippinglabel.de/v2"

type Client struct {
	baseURL      string
	clientID     string
	clientSecret string
	hc           *http.Client
}

func NewClient(clientID string, clientSecret string) (*Client, error) {
	if clientID == "" || clientSecret == "" {
		return nil, types.ErrRequiredClientIDAndSecret
	}

	c := &Client{clientID: clientID, clientSecret: clientSecret}
	c.Development()
	c.hc = c.defaultHTTPClient()
	return c, nil
}

// Production sets the productionURL as default
func (c *Client) Production() {
	c.baseURL = productionURL
}

// Development sets the developmentURL as default
func (c *Client) Development() {
	c.baseURL = developmentURL
}

// SetHTTPClient sets the default http.Client
func (c *Client) SetHTTPClient(hc *http.Client) {
	c.hc = hc
}

// defaultHTTPClient sets the default http.Client
func (c *Client) defaultHTTPClient() *http.Client {
	t := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	_ = http2.ConfigureTransport(t)

	hc := &http.Client{
		Transport: t,
		Timeout:   2 * time.Minute,
	}
	return hc
}

// APIContext creates a token specific context for the REST API
func (c *Client) APIContext(token *types.AuthToken) (*APIContext, error) {
	return NewAPIContext(c, token)
}

// send sends the request to the Shippinglabel REST API
func (c *Client) send(ctx context.Context, req *request) error {
	httpReq, err := req.HTTPRequest(ctx)
	if err != nil {
		return err
	}

	// Send request
	resp, err := c.hc.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode >= 400 {
		em := &types.Error{}
		if err = json.NewDecoder(resp.Body).Decode(em); err != nil {
			return err
		}
		return em
	}

	// Check resp handler was set
	if req.respHandler == nil {
		_, err = io.ReadAll(resp.Body)
		return err
	}

	// Handle response
	return req.respHandler(resp)
}

// sendTokenRequest sends a request to receive an access token for the ClientCredentials, AuthorizationCode and RefreshToken functions
func (c *Client) sendTokenRequest(ctx context.Context, qs url.Values) (*types.AuthToken, error) {
	var tk *types.AuthToken
	req := newRequest(c.baseURL).SetBasicAuth(c.clientID, c.clientSecret).SetFormValues(qs).ToJSON(&tk).
		SetMethod(http.MethodPost).SetPath("/oauth2/token")
	if err := c.send(ctx, req); err != nil {
		return nil, err
	}
	tk.SetExpirationTime()
	return tk, nil
}

// ClientCredentials creates an access token with the client credentials
func (c *Client) ClientCredentials(ctx context.Context) (*types.AuthToken, error) {
	qs := url.Values{}
	qs.Add("grant_type", "client_credentials")
	return c.sendTokenRequest(ctx, qs)
}

// AuthCodeURL creates a redirect url for the shippinglabel oauth process (AuthorizationCode)
func (c *Client) AuthCodeURL(redirectURL string, state string) string {
	qs := url.Values{}
	qs.Add("redirect_uri", redirectURL)
	qs.Add("response_type", "code")
	if state != "" {
		qs.Add("state", state)
	}
	return fmt.Sprintf("%s/oauth2/authorize?%s", c.baseURL, qs.Encode())
}

// AuthorizationCode exchanges the authorization code for an access token
func (c *Client) AuthorizationCode(ctx context.Context, authCode string) (resp *types.AuthToken, err error) {
	qs := url.Values{}
	qs.Add("grant_type", "authorization_code")
	qs.Add("code", authCode)
	return c.sendTokenRequest(ctx, qs)
}

// RefreshToken creates an access token through a refresh token
func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (resp *types.AuthToken, err error) {
	qs := url.Values{}
	qs.Add("grant_type", "refresh_token")
	qs.Add("refresh_token", refreshToken)
	return c.sendTokenRequest(ctx, qs)
}
