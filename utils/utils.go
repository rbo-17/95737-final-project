package utils

import "errors"

func TestTypeToWriteFactor(testType TestType) (float64, error) {
	var writeFactor float64
	switch testType {
	case TestTypeRead:
		writeFactor = WriteFactorRead
	case TestTypeBalanced:
		writeFactor = WriteFactorBalanced
	case TestTypeWrite:
		writeFactor = WriteFactorWrite
	default:
		return 0, errors.New("invalid test data type provided")
	}

	return writeFactor, nil
}
