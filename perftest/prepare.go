package perftest

import (
	"fmt"
	"github.com/rbo-17/95737-final-project/setup"
	"github.com/rbo-17/95737-final-project/utils"
	"time"
)

type TestOp struct {
	KeyId     int64
	Value     []byte
	ValueSize int
	OpType    utils.OpType
}

type TestOpResult struct {
	Time    time.Time
	OpType  utils.OpType
	KeyId   int64
	Latency int64
	Ok      bool
	Err     error
}

func GetTestOp(existingRecords, newRecords []setup.TestRecord, newRecordI *int, op utils.OpType) TestOp {

	var key int64
	var payload []byte
	var payloadSize int

	switch op {
	case utils.OpTypeGet:

		// Get record from starter dataset
		keyI := setup.GetRandLimitInt64(len(existingRecords))
		key = existingRecords[keyI].KeyId
		payload = nil
		payloadSize = existingRecords[keyI].ValueSize

	case utils.OpTypePut:

		// Get record from new dataset
		newRecord := newRecords[*newRecordI]

		key = newRecord.KeyId
		payload = newRecord.Value
		payloadSize = newRecord.ValueSize

		*newRecordI++
	}

	return TestOp{
		KeyId:     key,
		Value:     payload,
		ValueSize: payloadSize,
		OpType:    op,
	}
}

func GetTestOps(existingRecords, newRecords []setup.TestRecord, writeFactor float64) []TestOp {

	// Total op count will equal existing records count
	opCount := len(existingRecords)
	ops := make([]TestOp, opCount)

	var newRecordI int
	var rCnt, wCnt int
	for i := 0; i < opCount; i++ {

		// Get operation
		opBool := setup.GetRandBoolWeighted(writeFactor)
		var op utils.OpType
		if opBool {
			op = utils.OpTypePut
			wCnt++
		} else {
			op = utils.OpTypeGet
			rCnt++
		}

		t := GetTestOp(existingRecords, newRecords, &newRecordI, op)
		ops[i] = t
	}

	utils.Print(fmt.Sprintf("Created %d read ops and %d write ops", rCnt, wCnt))

	return ops
}

func Prepare(starterRecordsMap, newRecordsMap map[int64]setup.TestRecord, testType utils.TestType, dataType utils.TestDataType) ([]TestOp, error) {

	// Convert maps to lists
	starterRecordsList := setup.RecordsMapToList(starterRecordsMap)
	newRecordsList := setup.RecordsMapToList(newRecordsMap)

	// Determine how many ops should be reads. Value will be between 0 and 1.
	writeFactor, err := utils.TestTypeToWriteFactor(testType)
	if err != nil {
		return nil, err
	}

	// Build test operations (ops)
	ops := GetTestOps(starterRecordsList, newRecordsList, writeFactor)

	return ops, nil
}
