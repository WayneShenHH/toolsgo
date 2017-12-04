package repositoryimpl

// Rpush call redis Rpush
func (db *datastore) Rpush(key string, value []byte) {
	_, err := db.cache.Db.Do("rpush", key, value)
	if err != nil {
		panic(err)
	}
}

// Hset call redis Hset
func (db *datastore) Hset(key string, field string, value []byte) {
	_, err := db.cache.Db.Do("hset", key, field, value)
	if err != nil {
		panic(err)
	}
}
