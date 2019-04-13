package backend_api

type Config struct {
	Port string `toml:"apiPort"`
	FeatherBaseURL string

	// Weather service
	WeatherBaseUrl string `toml:"weatherBaseUrl"`// contains username + secret

}
