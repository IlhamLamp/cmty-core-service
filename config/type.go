package config

type Config struct {
	AppClient           string
	AppPort             string
	AppName             string
	CloudinaryApiKey    string
	CloudinaryApiSecret string
	CloudinaryCloudName string
	JwtAccessSecret     string
	JwtRefreshSecret    string
}
