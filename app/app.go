// Package app includes business logic for afu-entertainment.
// It also serves as a module including shared functions.
package app

import (
	"os"

	"github.com/spf13/viper"
)

// JWTSecret 產生憑證密鑰
const (
	JWTSecret  = "123"
	LessAmount = 50 //投注最低限額設定 (人民幣)
)

// General setting
const (
	UserOfferLimit = -500000 // 單使用者 Offer 最多可贏金額
	BankLockLimit  = 1       // 出金失敗到達此次數時鎖定帳號
)

// WorkerKey worker key in redis
type WorkerKey struct {
	Timestamp string
	Duration  string
	Message   string
	Counter   string
}

// WorkerKeys redis key
type WorkerKeys struct {
	Cache         WorkerKey
	Match         WorkerKey
	Offer         WorkerKey
	OfferSettle   WorkerKey
	Comparer      WorkerKey
	Variant       WorkerKey
	Log           WorkerKey
	ScoreResult   string
	Settle        string
	PushRawResult string
	MCT           WorkerKey
	MatchStat     WorkerKey
	Pevt          WorkerKey
	Resulting     WorkerKey
	LiveScore     WorkerKey
	DeletedPeid   WorkerKey
	//using for websocket
	PlayerChannel   WorkerKey
	OperatorChannel WorkerKey
	PositionWarning WorkerKey
	PlayerOrder     WorkerKey
	OperatorOrder   WorkerKey
}

// Setting 全系統設定吃這個
var Setting EnvironmentConfig
var appkeys *WorkerKeys

// BuildKeys build redis key
func BuildKeys() WorkerKeys {
	if appkeys == nil {
		appkeys = &WorkerKeys{
			Comparer:        buildWorkerKey("match"),
			Offer:           buildWorkerKey("offer"),
			OfferSettle:     buildWorkerKey("offersettle"),
			Variant:         buildWorkerKey("variant"),
			Log:             buildWorkerKey("push_log"),
			Cache:           buildWorkerKey("cache"),
			Match:           buildWorkerKey("push_match"),
			PlayerChannel:   buildWorkerKey("player_channel"),
			OperatorChannel: buildWorkerKey("operator_channel"),
			PositionWarning: buildWorkerKey("position_warning"),
			PlayerOrder:     buildWorkerKey("player_order"),
			OperatorOrder:   buildWorkerKey("operator_order"),
			ScoreResult:     "worker:result:score",
			Settle:          "worker:settle:data",
			PushRawResult:   "worker:push:result",
			MCT:             buildWorkerKey("mct"),
			MatchStat:       buildWorkerKey("matchstat"),
			Pevt:            buildWorkerKey("pevt"),
			LiveScore:       buildWorkerKey("livescore"),
			Resulting:       buildWorkerKey("resulting"),
			DeletedPeid:     buildWorkerKey("deletedpeid"),
		}
	}
	return *appkeys
}

func buildWorkerKey(worker string) WorkerKey {
	return WorkerKey{
		Timestamp: "worker:" + worker + ":timestamp",
		Duration:  "worker:" + worker + ":processing_time",
		Message:   "worker:" + worker + ":message",
		Counter:   "worker:" + worker + ":counter",
	}
}

// Env get env setting
func Env() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development"
	}
	return env
}

// Config get config by env setting
func Config() {

	var c EnvironmentConfig
	viper.Unmarshal(&c)
	Setting = c
}

// EnvironmentConfig config model of env
type EnvironmentConfig struct {
	Database  *DatabaseConfig
	HTTP      *HTTPConfig
	Redis     *RedisConfig
	Notify    *NotifyConfig
	Websocket *WebSocketConfig
	Swagger   *SwaggerConfig
	Logger    *LoggerConfig
}

// DatabaseConfig 資料庫連線設定
type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	Encoding     string
	Name         string
	Timeout      int
	MaxIdleConns int
	MaxConns     int
	LogMode      bool
}

// LoggerConfig logger setting
type LoggerConfig struct {
	StdLevel  string
	FileLevel string
}

// SwaggerConfig swagger 相關設定
type SwaggerConfig struct {
	Enable   bool
	FilePath string
}

// HTTPConfig http setting
type HTTPConfig struct {
	Addr        string
	PingAddr    string
	B2BAddr     string
	PingB2BAddr string
	WSAddr      string
	PingWSAddr  string
}

// WebSocketConfig web socket setting
type WebSocketConfig struct {
	PingDelay int
}

// RedisConfig redis 設定
type RedisConfig struct {
	Host         string
	Port         int
	Index        int
	MaxIdleConns int
	MaxConns     int
}

// NotifyConfig 設定通知
type NotifyConfig struct {
	SlackEnable bool
}

func readConfig(env string) EnvironmentConfig {
	var config EnvironmentConfig
	return config
}
