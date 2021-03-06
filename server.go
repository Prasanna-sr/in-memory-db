package main

import "os"
import "fmt"
import "bufio"
import "strings"

import "in-memory-db/dbtransaction"

var mdb = dbtransaction.NewDb()

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	for text != "END\n" {
		handleDbCommands(strings.Trim(text, "\n"))
		text, _ = reader.ReadString('\n')
	}
}

func handleDbCommands(statement string) {
	cmdList := strings.Split(statement, " ")
	command := cmdList[0]
	if command == "BEGIN" {
		mdb.StartTransaction()
	}
	if command == "ROLLBACK" {
		if !mdb.Rollback() {
			fmt.Println("NO TRANSACTION")
		}
	}
	if command == "COMMIT" {
		if !mdb.StopAllTransaction() {
			fmt.Println("NO TRANSACTION")
		}
	}
	if command == "SET" {
		key := cmdList[1]
		value := cmdList[2]
		mdb.Set(key, value)
	}
	if command == "GET" {
		key := cmdList[1]
		if mdb.Get(key) != "" {
			fmt.Println(mdb.Get(key))
		} else {
			fmt.Println("NULL")
		}
	}
	if command == "UNSET" {
		key := cmdList[1]
		mdb.Unset(key)
	}
	if command == "NUMEQUALTO" {
		value := cmdList[1]
		count := mdb.NumCount(value)
		fmt.Println(count)
	}

}
