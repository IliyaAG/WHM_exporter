package whm
import (
    "encoding/json"
    "os/exec"
)

type Client struct {
    timeout int
}

func NewClient(cfg Config) *Client {
    return &Client{timeout: cfg.Timeout}
}

func (c *Client) ListAccounts() (*ListAccountsResponse, error) {
    cmd := exec.Command("whmapi1", "--output=json", "listaccts")

    out, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    var resp ListAccountsResponse
    if err := json.Unmarshal(out, &resp); err != nil {
        return nil, err
    }

    return &resp, nil
}
