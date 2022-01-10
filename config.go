package biteship

const (
	DEFAULT_BITESHIP_URL = "https://api.biteship.com"
)

type ConfigOption struct {
	SecretKey   string
	BiteshipUrl string
}

func DefaultConfig(secretKey string) *ConfigOption {
	return &ConfigOption{
		SecretKey:   secretKey,
		BiteshipUrl: DEFAULT_BITESHIP_URL,
	}
}
