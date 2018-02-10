package mysql

import (
	"fmt"

	"github.com/hlts2/go-check-database-tables/dao/databases/config"
)

type mysqlConfig struct {
	config.Config
}

func (c mysqlConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", c.User, c.Password, c.Host, c.Database)
}

func (c mysqlConfig) DatabaseName() string {
	return c.Database
}

//GetMysqlConfig creates mysqlConfig instance
func GetMysqlConfig(c config.Config) config.DBConfig {
	return mysqlConfig{c}
}
