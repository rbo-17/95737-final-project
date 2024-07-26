package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rbo-17/95737-final-project/perftest"
	"github.com/rbo-17/95737-final-project/setup"
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
	} else if len(os.Args) < 3 {
		setup.PrintHelpTestType()
	} else if len(os.Args) < 4 {
		setup.PrintHelpDataType()
	}

	dbNameArg := os.Args[1]
	testTypeArg := os.Args[2]
	dataTypeArg := os.Args[3]

	db := setup.ValidateDbNameArg(dbNameArg)
	testType := setup.ValidateTestTypeArg(testTypeArg)
	dataType := setup.ValidateDataTypeArg(dataTypeArg)

	perftest.RunTest(db, testType, dataType)
}
