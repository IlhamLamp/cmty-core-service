package types

type Config struct {
	AppClient           string
	AppPort             string
	AppName             string
	DbUsername          string
	DbPassword          string
	DbName              string
	DbHost              string
	CloudinaryApiKey    string
	CloudinaryApiSecret string
	CloudinaryCloudName string
	JwtAccessSecret     string
	JwtRefreshSecret    string
}
