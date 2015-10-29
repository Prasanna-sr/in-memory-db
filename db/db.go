package db

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
	_, ok := memory.m[key]
	if ok == true {
		oldValue := memory.m[key]
		memory.c[oldValue] = memory.c[oldValue] - 1
	}
	memory.m[key] = value
	memory.c[value] = memory.c[value] + 1

}

func (memory *Memorydb) Unset(key string) {
	value, ok := memory.m[key]
	if ok == true {
		memory.c[value] = memory.c[value] - 1
	}
	delete(memory.m, key)
}

func (memory Memorydb) NumCount(value string) int {
	return memory.c[value]
}
