package dbtransaction

import "in-memory-db/db"

type dbTMap struct {
	m map[string]string
	c map[string]int
}
type dbTList []dbTMap

type TranDB struct {
	tran bool
	dbL  dbTList
	dbM  dbTMap
}

var mdb = db.NewDb()

func NewDb() TranDB {
	tDb := TranDB{}
	tDb.dbM.m = make(map[string]string)
	tDb.dbM.c = make(map[string]int)
	return tDb
}

func (t *TranDB) StartTransaction() {
	if t.tran == false {
		t.tran = true
		dbM := dbTMap{make(map[string]string), make(map[string]int)}
		t.dbL = append(t.dbL, dbM)
	} else {
		newDbM := dbTMap{make(map[string]string), make(map[string]int)}
		for k, v := range t.dbM.m {
			newDbM.m[k] = v
		}
		for k, v := range t.dbM.c {
			newDbM.c[k] = v
		}
		t.dbL = append(t.dbL, newDbM)
	}

}

func (t *TranDB) Rollback() bool {
	if t.tran == true {
		t.dbL = t.dbL[1:]
		if len(t.dbL) == 0 {
			t.tran = false
			t.dbM = dbTMap{make(map[string]string), make(map[string]int)}
		} else {
			t.dbM = dbTMap{make(map[string]string), make(map[string]int)}
			for k, v := range t.dbL[0].m {
				t.dbM.m[k] = v
			}
			for k, v := range t.dbL[0].c {
				t.dbM.c[k] = v
			}
		}
		return true
	}
	return false
}

func (t *TranDB) StopAllTransaction() bool {
	if t.tran == true {
		t.tran = false
		for k, v := range t.dbM.m {

			if v == "-1" {
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

func (t *TranDB) Get(key string) string {
	if t.tran == true {
		elem, ok := t.dbM.m[key]
		if ok == true {
			if elem == "-1" {
				return ""
			}
			return elem
		}
	}
	return mdb.Get(key)
}

func (t *TranDB) Set(key string, value string) {
	if t.tran == true {

		_, ok1 := t.dbM.m[key]
		if ok1 == true {
			oldValue := t.dbM.m[key]
			t.dbM.c[oldValue] = t.dbM.c[oldValue] - 1
		}
		t.dbM.m[key] = value
		_, ok := t.dbM.c[value]
		if ok == true {
			t.dbM.c[value] = t.dbM.c[value] + 1
		} else {
			t.dbM.c[value] = mdb.NumCount(value) + 1
		}
	} else {
		mdb.Set(key, value)
	}
}

func (t *TranDB) Unset(key string) {
	if t.tran == true {

		value, ok := t.dbM.m[key]
		if ok == false {
			value = mdb.Get(key)
		}
		t.dbM.m[key] = "-1"

		_, ok1 := t.dbM.c[value]
		if ok1 == true {
			t.dbM.c[value] = t.dbM.c[value] - 1
		} else {
			t.dbM.c[value] = mdb.NumCount(value) - 1
		}
	} else {
		mdb.Unset(key)
	}

}

func (t *TranDB) NumCount(value string) int {
	if t.tran == true {
		_, ok := t.dbM.c[value]
		if ok == true {
			return t.dbM.c[value]
		} else {
			return mdb.NumCount(value)
		}

	} else {
		return mdb.NumCount(value)
	}

}
