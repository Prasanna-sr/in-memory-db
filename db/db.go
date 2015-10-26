package db

// import "fmt"

type Memorydb struct {
}

var m = make(map[string]int)
var c = make(map[int]int)

func NewDb() Memorydb {
	mdb := Memorydb{}
	return mdb
}
func (memory Memorydb) Get(key string) int {
	return m[key]
}

func (memory Memorydb) Set(key string, value int) {
	m[key] = value
	c[value] = c[value] + 1
}

func (memory Memorydb) Unset(key string) {
	value := m[key]
	if c[value] > 0 {
		c[value] = c[value] - 1
	}
	delete(m, key)
}

func (memory Memorydb) NumCount(value int) int {
	return c[value]
}
