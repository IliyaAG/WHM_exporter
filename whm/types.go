package whm
type ListAccountsResponse struct {
    Metadata Metadata `json:"metadata"`
    Data     Data     `json:"data"`
}

type Metadata struct {
    Result int `json:"result"`
}

type Data struct {
    Accounts []Account `json:"acct"`
}

type Account struct {
    User          string `json:"user"`
    Domain        string `json:"domain"`
    Plan          string `json:"plan"`
    DiskLimit     string `json:"disklimit"`
    DiskUsed      string `json:"diskused"`
    Suspended     int    `json:"suspended"`
    SuspendReason string `json:"suspendreason"`
    MaxAddon      string `json:"maxaddon"`
    MaxParked     string `json:"maxparked"`
    MaxSub        string `json:"maxsub"`
    Owner         string `json:"owner"`
}
