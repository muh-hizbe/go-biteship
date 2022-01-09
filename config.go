package biteship

const (
	DEFAULT_BITESHIP_URL = "https://api.biteship.com"
)

type ConfigOption struct {
	SecretKey   string
	BiteshipUrl string
}

//var Config ConfigOption = ConfigOption{
//	BiteshipUrl: DEFAULT_BITESHIP_URL,
//}

func DefaultConfig(secretKey string) *ConfigOption {
	return &ConfigOption{
		SecretKey:   secretKey,
		BiteshipUrl: DEFAULT_BITESHIP_URL,
	}
}
