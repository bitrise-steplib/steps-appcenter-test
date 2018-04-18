package appcenter

// Client ...
type Client struct {
	apiToken string
}

// NewClient ...
func NewClient(token string) Client {
	return Client{
		apiToken: token,
	}
}
