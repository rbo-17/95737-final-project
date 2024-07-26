package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rbo-17/95737-final-project/perftest"
	"github.com/rbo-17/95737-final-project/setup"
	"github.com/rbo-17/95737-final-project/utils"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file!")
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) < 2 {
		setup.PrintHelpDbName()
	}

	dbNameArg := os.Args[1]
	db := setup.ValidateDbNameArg(dbNameArg)

	testTypes := []utils.TestType{utils.TestTypeRead, utils.TestTypeBalanced, utils.TestTypeWrite}
	dataTypes := []utils.TestDataType{utils.TestDataTypeSm, utils.TestDataTypeLg, utils.TestDataTypeImg}

	for _, testType := range testTypes {
		for _, dataType := range dataTypes {
			perftest.RunTest(db, testType, dataType)
		}
	}
}
