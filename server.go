package main

import "os"
import "fmt"
import "bufio"
import "strings"

import "strconv"

import "in-memory-db/db"

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	mdb := db.NewDb()
	for text != "END\n" {
		handleDbCommands(mdb, strings.Trim(text, "\n"))
		text, _ = reader.ReadString('\n')
	}
}

func handleDbCommands(mdb db.Memorydb, statement string) {
	cmdList := strings.Split(statement, " ")
	command = cmdList[0]
	key := cmdList[1]
	if command == "SET" {
		value, err := strconv.Atoi(cmdList[2])
		if err != nil {
			fmt.Println(err)
		}
		mdb.Set(key, value)
	}
	if command == "GET" {
		fmt.Println(mdb.Get(key))
	}
	if command == "UNSET" {
		mdb.Unset(key)
	}
	if command == "NUMEQUALTO" {
		fmt.Println(mdb.NumCount(key))
	}

}
