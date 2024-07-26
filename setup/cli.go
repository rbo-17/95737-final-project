package setup

import (
	"fmt"
	dbi "github.com/rbo-17/95737-final-project/db"
	//dbmongodb "github.com/rbo-17/95737-final-project/db/mongodb"
	dbmysql "github.com/rbo-17/95737-final-project/db/mysql"
	dbredis "github.com/rbo-17/95737-final-project/db/redis"
	"github.com/rbo-17/95737-final-project/utils"
	"os"
	"strings"
)

func PrintHelpDbName() {
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

func PrintHelpTestType() {
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

func PrintHelpDataType() {
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

func ValidateDbNameArg(dbNameArg string) dbi.Db {

	// Validate dbName arg and return db to benchmark
	var db dbi.Db
	switch dbNameArg {
	case utils.DbNameRedis:
		db = dbredis.NewRedis()

	case utils.DbNameMongoDB:
		panic("not implemented yet!")
		//db = dbmongodb.NewMongoDB()

	case utils.DbNameCassandra:
		panic("not implemented yet!")

	case utils.DbNameMySQL:
		db = dbmysql.NewMySQL()

	default:
		PrintHelpDbName()
	}

	return db
}

func ValidateTestTypeArg(testTypeArg string) utils.TestType {

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
		PrintHelpTestType()
	}

	return testType

}

func ValidateDataTypeArg(dataTypeArg string) utils.TestDataType {

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
		PrintHelpDataType()
	}

	return dataType
}
