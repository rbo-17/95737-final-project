package setup

import (
	"github.com/rbo-17/95737-final-project/utils"
	"io"
	"os"
)

const TextPath = "data/the-republic.txt"

const smallTextLenMin = 100
const smallTextLenMax = 200

const largeTextLenMin = 10000
const largeTextLenMax = 20000

var textStr string
var textLen int
var textLenStartMax int

func init() {

	// Load text
	file, err := os.Open(TextPath)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	textStr = string(b)
	textLen = len(textStr)

}

func GetNextText(start, end int) string {

	textLenStartMax = textLen - end

	startChar := GetRandLimit(textLenStartMax)      // Get random start
	endChar := startChar + GetRandRange(start, end) // Get random end between start & end

	return textStr[startChar:endChar]
}

func GetNextSmallText(opts utils.TestOpts) string {
	return GetNextText(smallTextLenMin*opts.DenormalizationFactor, smallTextLenMax*opts.DenormalizationFactor)
}

func GetNextSmallTextBytes(opts utils.TestOpts) ([]byte, error) {
	return []byte(GetNextSmallText(opts)), nil
}

func GetNextLargeText(opts utils.TestOpts) string {
	return GetNextText(largeTextLenMin*opts.DenormalizationFactor, largeTextLenMax*opts.DenormalizationFactor)
}

func GetNextLargeTextBytes(opts utils.TestOpts) ([]byte, error) {
	return []byte(GetNextLargeText(opts)), nil
}
