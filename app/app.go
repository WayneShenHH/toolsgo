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

var appkeys *Keys

type Keys struct {
	BroadcastPlayerChannel   string
	BroadcastOperatorChannel string
}

func GetKeys() *Keys {
	if appkeys == nil {
		appkeys = &Keys{
			BroadcastPlayerChannel:   "Broadcast:Player",
			BroadcastOperatorChannel: "Broadcast:Operator",
		}
	}
	return appkeys
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
	Mysql              string
	Redis              string
	RedisDatabaseIndex int
}

func readConfig(env string) EnvironmentConfig {
	var config EnvironmentConfig
	switch env {
	case "development":
		config = EnvironmentConfig{
			Mysql:              "root:123456@tcp(localhost:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis:              ":6379",
			RedisDatabaseIndex: 0,
		}
	case "test":
		config = EnvironmentConfig{
			Mysql:              "root:123456@tcp(localhost:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis:              ":6379",
			RedisDatabaseIndex: 0,
		}
	case "production":
		config = EnvironmentConfig{
			Mysql:              "root:123456@tcp(localhost:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis:              ":6379",
			RedisDatabaseIndex: 0,
		}
	case "remote":
		config = EnvironmentConfig{
			Mysql:              "root:123456@tcp(localhost:3306)/sbodds?charset=utf8&parseTime=true&loc=UTC&multiStatements=true",
			Redis:              ":6979",
			RedisDatabaseIndex: 1,
		}
	}

	return config
}
