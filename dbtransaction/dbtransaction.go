package dbtransaction

import "in-memory-db/db"

import "fmt"

type Memorydb struct {
	tran   bool
	dbList dbTList
}
type dbTran struct {
	m map[string]int
	c map[int]int
}
type dbTList []dbTran

var mdb = db.NewDb()

var dbList dbTList
var dbT dbTran

func NewDb() Memorydb {
	db := Memorydb{}
	return db
}

func (memory *Memorydb) StartTransaction() {
	if memory.tran == false {
		memory.tran = true
		dbT = dbTran{make(map[string]int), make(map[int]int)}
		dbList = append(dbList, dbT)
	} else {
		newDbt := dbTran{make(map[string]int), make(map[int]int)}
		newDbt.m = dbT.m
		for k, v := range dbT.m {
			newDbt.m[k] = v
		}
		for k, v := range dbT.c {
			newDbt.c[k] = v
		}
		// newDbt.c = dbT.c
		dbList = append(dbList, newDbt)
	}

}
func (memory Memorydb) Rollback() string {
	if memory.tran == true {
		fmt.Println(dbList)
		dbList = dbList[:len(dbList)-1]
		if len(dbList) == 0 {
			memory.tran = false
			dbT = dbTran{make(map[string]int), make(map[int]int)}
		} else {
			dbT.m = dbList[len(dbList)-1].m
			dbT.c = dbList[len(dbList)-1].c
			fmt.Println(dbT)
		}
	} else {
		return "NO TRANSACTION"
	}
	return ""
}
func (memory Memorydb) StopAllTransaction() bool {
	memory.tran = false
	dbList = dbList[:0]
	return false
}

func (memory Memorydb) Get(key string) int {
	if memory.tran == true {
		elem, ok := dbT.m[key]
		if ok == true {
			return elem
		}
	}
	return mdb.Get(key)

}

func (memory Memorydb) Set(key string, value int) {
	if memory.tran == true {
		dbList[len(dbList)-1].c[value] = mdb.NumCount(value) + 1
		dbList[len(dbList)-1].m[key] = value
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
