package questionbank

import "sync"

var (
	once sync.Once
	key1 = false
	key2 = false
)
