package config

import "sync"

// Config ...
type Config struct {
	MySQLConnectionString string
}

var pkgInit sync.Once
var pkgConfig *Config

// GetConfig ...
func GetConfig() *Config {
	pkgInit.Do(func() {
		pkgConfig = &Config{
			MySQLConnectionString: "root:@/hotel?charset=utf8&parseTime=True&loc=Local",
		}
	})

	return pkgConfig
}
