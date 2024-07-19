package utils

import (
	"fmt"
	"log"
)

//
//import (
//	"log"
//	"os"
//)
//
//func init(dbName string) {
//
//}
//
//func GetLogger() {
//	logger := log.New(os.Stdout, dbName, log.LstdFlags)
//
//	logger.SetPrefix()
//}

type LogPrefixData struct {
	DbName   string
	TestType TestType
	DataType TestDataType
}

var l LogPrefixData

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func UpdatePrefix(dbName string, testType TestType, dataType TestDataType) {
	l.DbName = dbName
	l.TestType = testType
	l.DataType = dataType

	prefix := fmt.Sprintf("[%s][%s][%s] ", l.DbName, l.TestType, l.DataType)
	log.SetPrefix(prefix)
}

func Print(msg string) {
	if l.DbName == "" || l.TestType == "" || l.DataType == "" {
		log.Println("[ERROR] Logger not initialized!")
		return
	}

	log.Println(msg)
}
