package setup

import (
	"github.com/rbo-17/95737-final-project/utils"
	"math"
)

type TestRecord struct {
	KeyId     int64
	Value     []byte
	ValueSize int
}

type TestDataSet struct {
	SmallTexts map[int64]TestRecord
	LargeTexts map[int64]TestRecord
	Images     map[int64]TestRecord
}

func RecordsMapToList(m map[int64]TestRecord) []TestRecord {

	l := make([]TestRecord, len(m))
	i := 0
	for k, v := range m {
		l[i] = TestRecord{
			KeyId:     k,
			Value:     v.Value,
			ValueSize: v.ValueSize,
		}
		i++
	}

	return l
}

type GetTestDataPayload func(opts utils.TestOpts) ([]byte, error)

func GetTestData(fn GetTestDataPayload, opts utils.TestOpts) (TestRecord, error) {

	payload, err := fn(opts)
	if err != nil {
		return TestRecord{}, err
	}

	return TestRecord{
		Value:     payload,
		ValueSize: len(payload),
	}, nil
}

func GetStarterDataSet(testType utils.TestDataType, opts utils.TestOpts) (map[int64]TestRecord, error) {
	return GetTestDataSet(1.0, testType, opts)
}

func GetTestDataSet(factor float64, dataType utils.TestDataType, opts utils.TestOpts) (map[int64]TestRecord, error) {

	// Set the base count of records (the count of records with denormalization factor 1)
	smTxtCnt := int(math.Pow(10, 7) * factor)
	lgTxtCnt := int(math.Pow(10, 5) * factor)
	imgCnt := int(math.Pow(10, 3) * factor)

	var cnt int
	var procFn GetTestDataPayload
	switch dataType {
	case utils.TestDataTypeSm:
		cnt = smTxtCnt
		procFn = GetNextSmallTextBytes
	case utils.TestDataTypeLg:
		cnt = lgTxtCnt
		procFn = GetNextLargeTextBytes
	case utils.TestDataTypeImg:
		cnt = imgCnt
		procFn = GetNextImageBytes
	}

	// Adjust the count based on the average record size (based on denormalization factor)
	cnt = cnt / opts.DenormalizationFactor

	dataSet := make(map[int64]TestRecord, cnt)
	for i := 0; i < cnt; i++ {

		// Create a new key/value test record.
		// Use a for loop in case random key has already been used.
		for {

			// Get key
			newKey := GetRandInt64()

			_, ok := dataSet[newKey]
			if ok {
				continue
			}

			// Get value
			data, err := GetTestData(procFn, opts)
			if err != nil {
				return nil, err
			}

			// Set both
			dataSet[newKey] = data
			break
		}
	}

	return dataSet, nil
}
