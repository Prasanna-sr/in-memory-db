package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "strconv"
import "in-memory-db/dbtransaction"

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	mdb := dbtransaction.NewDb()
	for text != "END\n" {
		handleDbCommands(mdb, strings.Trim(text, "\n"))
		text, _ = reader.ReadString('\n')
	}
}

func handleDbCommands(mdb dbtransaction.Memorydb, statement string) {
	cmdList := strings.Split(statement, " ")
	command := cmdList[0]
	if command == "BEGIN" {
		mdb.StartTransaction()
	}
	if command == "ROLLBACK" {
		mdb.Rollback()
	}
	if command == "BEGIN" {
		mdb.StopAllTransaction()
	}
	if command == "SET" {
		key := cmdList[1]
		value, err := strconv.Atoi(cmdList[2])
		if err != nil {
			fmt.Println(err)
		}
		mdb.Set(key, value)
	}
	if command == "GET" {
		key := cmdList[1]
		fmt.Println(mdb.Get(key))
	}
	if command == "UNSET" {
		key := cmdList[1]
		mdb.Unset(key)
	}
	if command == "NUMEQUALTO" {
		key := cmdList[1]
		value, _ := strconv.Atoi(key)
		count := mdb.NumCount(value)
		fmt.Println(count)
	}

}
