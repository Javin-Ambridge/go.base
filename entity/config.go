package entity

// Config is the servers config
type Config struct {
	// AppID is the applicationID
	AppID string `yaml:"appid"`

	// Env Specifies what env we are in, this is obtained from an env variable
	Env string

	// ServerPort specifies the server port
	ServerPort int `yaml:"serverPort"`
}
