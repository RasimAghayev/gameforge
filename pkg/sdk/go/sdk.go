// Go SDK for GameForge
package sdk

type Client struct{}

func NewClient() *Client { return &Client{} }

func (c *Client) RegisterPlayer(tenantID, externalID string) (string, error) { return "", nil }
func (c *Client) AwardPoints(tenantID, playerID string, amount int64, reason string) error { return nil }
