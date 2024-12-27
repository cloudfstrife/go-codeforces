package client

const (
	CF_URL = "http://codeforces.com/api"
)

type Config struct {
	ApiKey    string
	ApiSecret string
	Lang      string
}
