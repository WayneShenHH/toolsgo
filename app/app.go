package app

import (
	"fmt"
	"os"
)

// Env get env setting
func Env() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

// Configuration get config by env setting
func Configuration() EnvironmentConfig {
	env := Env()
	if env == "" {
		fmt.Println("No Environment specified. Prepend GO_ENV=production to the command if running in production.")
		fmt.Println("Using development as default.")
		env = "development"
	}
	return readConfig(env)
}

// EnvironmentConfig config model of env
type EnvironmentConfig struct {
	Mysql string
	Redis string
}

func readConfig(env string) EnvironmentConfig {
	var config EnvironmentConfig
	switch env {
	case "development":
		config = EnvironmentConfig{
			Mysql: "root:123456@tcp(localhost:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis: ":6379",
		}
	case "test":
		config = EnvironmentConfig{
			Mysql: "root:123456@tcp(localhost:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis: ":6379",
		}
	case "production":
		config = EnvironmentConfig{
			Mysql: "ranbow_cc:A34sADkjS234FF8dfX23kS4jA8f@tcp(afu.cd75gda2paem.ap-northeast-1.rds.amazonaws.com:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis: "afu.ze67a8.0001.apne1.cache.amazonaws.com:6379",
		}
	}
	return config
}
