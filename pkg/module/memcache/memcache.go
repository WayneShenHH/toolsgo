// Package memcache 處理快取資料存取資源
package memcache

// MemCache memory cache
type MemCache interface {
	KEYS(pattern string) ([]string, error)
	SET(key interface{}, value interface{}, options ...interface{}) (reply interface{}, err error)
	SETEX(key interface{}, value interface{}, expireSeconds uint) error
	SETToken(key interface{}, token interface{}) (err error)
	GET(key, model interface{}) error
	GETSTR(key interface{}) (string, error)
	MGET(keys []string) ([][]byte, error)
	DEL(key interface{}) (err error)
	DELs(key ...interface{}) (err error)
	HSET(key, field, value interface{}) (err error)
	HGET(key, field interface{}) (reply interface{}, err error)
	HDEL(key, field interface{}) (err error)
	EVAL(script string, options ...interface{}) (reply interface{}, err error)
	EXISTS(key interface{}) (bool, error)
	Flush() error
	RPush(key string, content interface{}) (err error)
}
