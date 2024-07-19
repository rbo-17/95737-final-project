package main

import (
	"fmt"
	"github.com/joho/godotenv"
	dbi "github.com/rbo-17/95737-final-project/db"
	dbredis "github.com/rbo-17/95737-final-project/db/redis"
	"github.com/rbo-17/95737-final-project/perftest"
	"github.com/rbo-17/95737-final-project/setup"
	"github.com/rbo-17/95737-final-project/utils"
	"os"
	"strings"
	"time"
)

func printHelpDbName() {
	options := []string{utils.DbNameRedis, utils.DbNameMongoDB, utils.DbNameCassandra, utils.DbNameMySQL}
	_, err := fmt.Fprintln(os.Stderr, "Invalid database provided! Please provide a valid option:")
	if err != nil {
		panic(err)
	}

	optsJoined := strings.Join(options, "\n  -")
	_, err = fmt.Fprintln(os.Stderr, "  -"+optsJoined)
	if err != nil {
		panic(err)
	}

	os.Exit(1)
}

func printHelpTestType() {
	options := []string{string(utils.TestTypeRead), string(utils.TestTypeBalanced), string(utils.TestTypeWrite)}
	_, err := fmt.Fprintln(os.Stderr, "Invalid test type provided! Please provide a valid option:")
	if err != nil {
		panic(err)
	}

	optsJoined := strings.Join(options, "\n  -")
	_, err = fmt.Fprintln(os.Stderr, "  -"+optsJoined)
	if err != nil {
		panic(err)
	}

	os.Exit(1)
}

func printHelpDataType() {
	options := []string{string(utils.TestDataTypeSm), string(utils.TestDataTypeLg), string(utils.TestDataTypeImg)}
	_, err := fmt.Fprintln(os.Stderr, "Invalid data type provided! Please provide a valid option:")
	if err != nil {
		panic(err)
	}

	optsJoined := strings.Join(options, "\n  -")
	_, err = fmt.Fprintln(os.Stderr, "  -"+optsJoined)
	if err != nil {
		panic(err)
	}

	os.Exit(1)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file!")
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) < 2 {
		printHelpDbName()
	} else if len(os.Args) < 3 {
		printHelpTestType()
	} else if len(os.Args) < 4 {
		printHelpDataType()
	}

	dbNameArg := os.Args[1]
	testTypeArg := os.Args[2]
	dataTypeArg := os.Args[3]

	// Validate dbName arg and get db to benchmark
	var db dbi.Db
	switch dbNameArg {
	case utils.DbNameRedis:
		db = dbredis.NewRedis()

	case utils.DbNameMongoDB:
		panic("not implemented yet!")

	case utils.DbNameCassandra:
		panic("not implemented yet!")

	case utils.DbNameMySQL:
		panic("not implemented yet!")

	default:
		printHelpDbName()
	}

	// Validate testType
	var testType utils.TestType
	switch testTypeArg {
	case string(utils.TestTypeRead):
		testType = utils.TestTypeRead

	case string(utils.TestTypeBalanced):
		testType = utils.TestTypeBalanced

	case string(utils.TestTypeWrite):
		testType = utils.TestTypeWrite

	default:
		printHelpTestType()
	}

	// Validate dataType
	var dataType utils.TestDataType
	switch dataTypeArg {
	case string(utils.TestDataTypeSm):
		dataType = utils.TestDataTypeSm

	case string(utils.TestDataTypeLg):
		dataType = utils.TestDataTypeLg

	case string(utils.TestDataTypeImg):
		dataType = utils.TestDataTypeImg

	default:
		printHelpDataType()
	}

	utils.UpdatePrefix(db.GetName(), testType, dataType)

	// Set up db connection and load test data
	db.Init()
	sds, err := setup.LoadStarterDataset(db, dataType)
	if err != nil {
		panic(err)
	}

	// Get new (unloaded) records to perform test with
	writeFactor, err := utils.TestTypeToWriteFactor(testType)
	if err != nil {
		panic(err)
	}

	//fmt.Println("writeFactor", writeFactor)

	// Add extra to write factor to account for random variation
	nds, err := setup.GetTestDataSet(writeFactor+0.01, dataType)
	if err != nil {
		panic(err)
	}

	//fmt.Println("nds len", len(nds))

	utils.Print("Loading new dataset for testing...")

	// Prepare & run test
	ops, err := perftest.Prepare(sds, nds, testType, dataType)
	if err != nil {
		panic(err)
	}

	utils.Print("Loading complete. Starting test now.")
	start := time.Now()
	err = perftest.Run(db, testType, dataType, ops)
	if err != nil {
		panic(err)
	}

	dur := int(time.Since(start).Seconds())
	utils.Print(fmt.Sprintf("Testing completed in %d seconds.", dur))

	// Clean up db
	err = db.DeleteAll()
	if err != nil {
		panic(err)
	}
}
