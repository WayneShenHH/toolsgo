// Package environment env configs
package environment

import "github.com/spf13/viper"

// Setting database parameters
var Setting Conf

// TestSetting database connection setting for testing
var TestSetting = Conf{
	Database: &DatabaseConfig{
		Username: "root",
		Password: "123456",
		Host:     "localhost",
		Port:     3306,
		Name:     "ftodds",
		Encoding: "utf8",
	},
	Redis: &RedisConfig{
		Host:         "localhost",
		Port:         6379,
		Index:        0,
		MaxIdleConns: 100,
		MaxConns:     100,
	},
	Push: &PushConfig{
		InsertDelay:     1,
		InsertChunkSize: 10,
	},
}

// environment parameters
// const (
// 	PRODUCTION = "production"
// 	STAGING    = "staging"
// 	DEVELOP    = "develop"
// 	TEST       = "test"
// 	goenv      = "GO_ENV"
// )

// redis task key
const (
	ODDSKEY        = "worker:match:message"
	PUSHRESULTKEY  = "worker:push:result"
	PEVTKEY        = "worker:pevt:message"
	LIVESCOREKEY   = "worker:livescore:message"
	RESULTINGKEY   = "worker:resulting:message"
	DELETEDPEIDKEY = "worker:deletedpeid:message"
)

// Config config parameters
func Config() {
	var c Conf
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}

	// if ENV = os.Getenv(goenv); ENV == "" {
	// 	ENV = DEVELOP
	// }
	Setting = c
	// switch ENV {
	// case PRODUCTION:
	// 	Setting = c.Production
	// case STAGING:
	// 	Setting = c.Staging
	// case TEST:
	// 	Setting = c.Test
	// case DEVELOP:
	// 	Setting = c.Develop
	// default:
	// 	logger.Error("No Environment specified. Prepend GO_ENV=production to the command if running in production.")
	// 	os.Exit(1)
	// }
	// logger.Info(fmt.Sprintf("%+v", Setting))
	// logger.Info(fmt.Sprintf("Redis:  %+v", Setting.Redis))
	// logger.Info(fmt.Sprintf("TxPush: %+v", Setting.TxPush))
	// logger.Info(fmt.Sprintf("MQ:     %+v", Setting.MQ))
	// logger.Info(fmt.Sprintf("Test:   %+v", Setting.Test))
}

// GetDBConfig 分解 config 屬性讓 wire 在注入的時候可以直接取得細部設定
func GetDBConfig(config Conf) *DatabaseConfig {
	return config.Database
}

// GetRedisConfig 分解 config 屬性讓 wire 在注入的時候可以直接取得細部設定
func GetRedisConfig(config Conf) *RedisConfig {
	return config.Redis
}

// GetTestConfig 分解 config 屬性讓 wire 在注入的時候可以直接取得細部設定
func GetTestConfig(config Conf) *TestConfig {
	return config.Test
}
