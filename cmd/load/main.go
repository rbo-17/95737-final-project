package main

import (
	"fmt"
	"github.com/rbo-17/95737-final-project/db/mysql"
	"github.com/rbo-17/95737-final-project/setup"
)

const tableNameSmallText = "SMALL_TEXT"
const tableNameLargeText = "LARGE_TEXT"
const tableNameImage = "IMAGES"

func BuildInsertInto(table string, records map[int64]setup.TestRecord) (string, []interface{}) {
	sqlStr := fmt.Sprintf("INSERT INTO %s(key, payload) VALUES ", table)
	var sqlVals []interface{}

	for k, v := range records {
		sqlStr += "(?, ?),"
		vals = append(sqlVals, k, v)
	}

	//trim the last ,
	sqlStr = sqlStr[0 : len(sqlStr)-1]

	return sqlStr, sqlVals
}

func main() {
	dataSet := setup.GetTestDataSet()
	db, err := mysql.GetMysqlInstance()
	if err != nil {
		panic(err)
	}

	for k, v := range dataSet.SmallTexts {
		sqlStr, sqlVals := BuildInsertInto(tableNameSmallText)

		stmt, err := db.Prepare(sqlStr)
		if err != nil {
			panic(err)
		}

		res, err := stmt.Exec(sqlVals...)
		if err != nil {
			panic(err)
		}
	}

}
