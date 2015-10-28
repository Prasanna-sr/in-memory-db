package dbtransaction

import "in-memory-db/db"
import "fmt"

type Memorydb struct {
	tran bool
}

var mdb = db.NewDb()
var mTran = make(map[string]int)
var cTran = make(map[int]int)
var tran bool

type dbTran struct {
	m map[string]int
	c map[int]int
}

type dbTList []dbTran

var dbList dbTList
var dbT dbTran

func NewDb() Memorydb {
	mdb := Memorydb{}
	return mdb
}

func (memory Memorydb) StartTransaction() {
	memory.tran = true
	dbT = dbTran{make(map[string]int), make(map[int]int)}
}
func (memory Memorydb) Rollback() {
	dbList = dbList[:len(dbList)-1]
}
func (memory Memorydb) StopAllTransaction() bool {
	memory.tran = false
	dbList = dbList[:0]
	return false
}

func (memory Memorydb) Get(key string) int {
	fmt.Println(memory.tran)
	if memory.tran == true {
		return dbList[0].m[key]
	} else {
		return mdb.Get(key)
	}

}

func (memory Memorydb) Set(key string, value int) {
	fmt.Println(memory.tran)
	if memory.tran == true {
		m := make(map[string]int)
		c := make(map[int]int)
		c[value] = mdb.NumCount(value) + 1
		m[key] = value
		dbLocal := dbTran{m, c}
		dbList = append(dbList, dbLocal)
	} else {
		mdb.Set(key, value)
	}
}

func (memory Memorydb) Unset(key string) {
	mdb.Unset(key)
}

func (memory Memorydb) NumCount(value int) int {
	return mdb.NumCount(value)
}
