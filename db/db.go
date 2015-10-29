package db

// import "fmt"

type Memorydb struct {
	m map[string]string
	c map[string]int
}

func NewDb() Memorydb {
	mdb := Memorydb{make(map[string]string), make(map[string]int)}
	return mdb
}

func (memory Memorydb) Get(key string) string {
	return memory.m[key]
}

func (memory *Memorydb) Set(key string, value string) {
	memory.m[key] = value
	memory.c[value] = memory.c[value] + 1
}

func (memory *Memorydb) Unset(key string) {
	value := memory.m[key]
	if memory.c[value] > 0 {
		memory.c[value] = memory.c[value] - 1
	}
	delete(memory.m, key)
}

func (memory Memorydb) NumCount(value string) int {
	return memory.c[value]
}
