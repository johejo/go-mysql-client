package gomysqlclient

import "fmt"

type Config struct {
	Host     string
	Port     uint
	User     string
	Password string
	Database string
}

func (c *Config) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Password, c.Host, c.Port, c.Database)
}
