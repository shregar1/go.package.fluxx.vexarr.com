// Package fluxx is the official Go SDK for the Fluxx control plane API.
package fluxx

// Version is the current SDK version.
const Version = "0.1.0"

// Client is the entry point for the Fluxx API.
type Client struct {
	apiKey string
}

// New constructs a Client bound to the given API key.
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
