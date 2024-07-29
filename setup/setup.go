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

type GetTestDataPayload func() []byte

func GetTestData(fn GetTestDataPayload) TestRecord {

	payload := fn()

	return TestRecord{
		Value:     payload,
		ValueSize: len(payload),
	}
}

func GetStarterDataSet(testType utils.TestDataType) (map[int64]TestRecord, error) {
	return GetTestDataSet(1.0, testType)
}

func GetTestDataSet(factor float64, dataType utils.TestDataType) (map[int64]TestRecord, error) {

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

	dataSet := make(map[int64]TestRecord, cnt)
	for i := 0; i < cnt; i++ {
		for {
			newKey := GetRandInt64()

			_, ok := dataSet[newKey]
			if ok {
				continue
			}

			data := GetTestData(procFn)
			dataSet[newKey] = data
			break
		}
	}

	return dataSet, nil
}
