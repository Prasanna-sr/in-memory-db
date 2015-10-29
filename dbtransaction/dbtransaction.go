package dbtransaction

import "in-memory-db/db"

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

		for k, v := range dbT.m {
			newDbt.m[k] = v
		}
		for k, v := range dbT.c {
			newDbt.c[k] = v
		}
		dbList = append(dbList, newDbt)
	}

}

func (memory *Memorydb) Rollback() bool {
	if memory.tran == true {
		dbList = dbList[1:]
		if len(dbList) == 0 {
			memory.tran = false
			dbT = dbTran{make(map[string]int), make(map[int]int)}
		} else {
			dbT = dbTran{make(map[string]int), make(map[int]int)}
			for k, v := range dbList[0].m {
				dbT.m[k] = v
			}
			for k, v := range dbList[0].c {
				dbT.c[k] = v
			}
		}
		return true
	}
	return false
}

func (memory *Memorydb) StopAllTransaction() bool {
	if memory.tran == true {
		memory.tran = false
		for k, v := range dbT.m {

			if v == -1 {
				mdb.Unset(k)
			} else {
				mdb.Set(k, v)
			}
		}
		return true
	} else {
		return false
	}
}

func (memory Memorydb) Get(key string) int {
	if memory.tran == true {
		elem, ok := dbT.m[key]
		if ok == true {
			if elem == -1 {
				return 0
			}
			return elem
		}
	}
	return mdb.Get(key)
}

func (memory Memorydb) Set(key string, value int) {
	if memory.tran == true {
		dbT.m[key] = value
		_, ok := dbT.c[value]
		if ok == true {
			dbT.c[value] = dbT.c[value] + 1
		} else {
			dbT.c[value] = mdb.NumCount(value) + 1
		}
	} else {
		mdb.Set(key, value)
	}
}

func (memory Memorydb) Unset(key string) {
	if memory.tran == true {

		value, ok := dbT.m[key]
		if ok == false {
			value = mdb.Get(key)
		}
		dbT.m[key] = -1

		_, ok1 := dbT.c[value]
		if ok1 == true {
			dbT.c[value] = dbT.c[value] - 1
		} else {
			dbT.c[value] = mdb.NumCount(value) - 1
		}
	} else {
		mdb.Unset(key)
	}

}

func (memory Memorydb) NumCount(value int) int {
	if memory.tran == true {
		_, ok := dbT.c[value]
		if ok == true {
			return dbT.c[value]
		} else {
			return mdb.NumCount(value)
		}

	} else {
		return mdb.NumCount(value)
	}

}
