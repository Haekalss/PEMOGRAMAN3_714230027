package config

var allowedOrigins = []string{
	"http://localhost:5173",
	"https://example.com",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
