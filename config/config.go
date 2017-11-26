package config

import "fmt"

// Config 配置信息
type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
	Viewer   string
	Output   string
	Debug    bool
}

func (c Config) String() string {
	return fmt.Sprintf("Host:\t%s\nPort:\t%d\nUser:\t%s\nDB:\t%s\nViewer:\t%s\nOutput:\t%s\nDebug:\t%t\n",
		c.Host,
		c.Port,
		c.Username,
		c.DB,
		c.Viewer,
		c.Output,
		c.Debug,
	)
}
