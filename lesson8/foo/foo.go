package foo

type Config struct {
	Port string
	DiagPort string
	Name string
	ID string
}

func NewConfig() *Config {
	new := &Config {
		Port: "",
		DiagPort: "",
		Name: "",
		ID: "",
	}
	return new
}
