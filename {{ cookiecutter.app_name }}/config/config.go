package config

// Config struct, change on demand
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	LogLevel string
	Dsn      string // Whether DSN must be defined depends on the condition of the monitored object
}
