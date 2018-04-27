package repositoryimpl

// Rpush call redis Rpush
func (db *datastore) Rpush(key string, value []byte) {
	r := db.cache.Get()
	defer r.Close()
	_, err := r.Do("rpush", key, value)
	if err != nil {
		panic(err)
	}
}

// Hset call redis Hset
func (db *datastore) Hset(key string, field string, value []byte) {
	r := db.cache.Get()
	defer r.Close()
	_, err := r.Do("hset", key, field, value)
	if err != nil {
		panic(err)
	}
}

// Blpop call redis Blpop
func (db *datastore) Blpop(key string) []byte {
	r := db.cache.Get()
	defer r.Close()
	inter, err := r.Do("blpop", key, 0)
	if inter == nil {
		return make([]byte, 0)
	}
	// BLPOP會回傳陣列第一組為key 第二組才是資料
	message := inter.([]interface{})[1]
	if err != nil {
		panic(err)
	}
	bytes, _ := message.([]byte)
	return bytes
}

func (db *datastore) LRange(key string, start int, end int) []interface{} {
	r := db.cache.Get()
	defer r.Close()
	inter, _ := r.Do("LRANGE", key, start, end)
	message := inter.([]interface{})
	return message
}
func (db *datastore) FlushDB() {
	r := db.cache.Get()
	defer r.Close()
	r.Do("FLUSHDB")
}
