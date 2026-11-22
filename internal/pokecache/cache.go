package pokecache



type Cache struct {
	Data map[string][]byte
}

func (c *Cache) NewCache() {
	c.Data = make(map[string][]byte)
}

func (c Cache) Add(key string, val []byte) {
	c.Data[key] = val
}

func (c Cache) Get(key string) ([]byte, bool) {
	value, ok := c.Data[key]
	if ok {
		return  value, ok
	}
	return nil, ok

}
