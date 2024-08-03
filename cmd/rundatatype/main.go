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
	} else if len(os.Args) < 3 {
		setup.PrintHelpDataType()
	}

	dbNameArg := os.Args[1]
	dataTypeArg := os.Args[2]

	db := setup.ValidateDbNameArg(dbNameArg)
	dataType := setup.ValidateDataTypeArg(dataTypeArg)

	testTypes := []utils.TestType{
		utils.TestTypeRead,
		utils.TestTypeBalanced,
		utils.TestTypeWrite,
	}

	for _, testType := range testTypes {
		// Always consider DenormalizationFactor with this test type. TODO: Add optional DenormalizationFactor to all tests
		for i := 1; i <= utils.DenormalizationFactor; i++ {

			opts := utils.TestOpts{
				DenormalizationFactor: i,
			}

			perftest.RunTest(db, testType, dataType, opts)
		}
	}
}
