package config

// Config 配置信息
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
	Tables   []string
	Viewer   string
	Output   string
	Debug    bool
}
