package config

var allowedOrigins = []string{
	"http://localhost:3000",
	"https://example.com",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
