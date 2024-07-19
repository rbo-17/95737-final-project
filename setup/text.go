package setup

import (
	"io"
	"math/rand"
	//"math/rand/v2"
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

	//var source int64 = 95737
	//rand.Seed(source)

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

	startChar := rand.Intn(textLenStartMax)
	endChar := startChar + (rand.Intn(end-start) + start)

	return textStr[startChar:endChar]
}

func GetNextSmallText() string {
	//txt := GetNextText(smallTextLenMin, smallTextLenMax)
	//fmt.Println("txt", txt)
	//fmt.Println("len(txt)", len(txt))

	return GetNextText(smallTextLenMin, smallTextLenMax)
}

func GetNextSmallTextBytes() []byte {
	return []byte(GetNextSmallText())
}

func GetNextLargeText() string {
	//fmt.Println("fetching text for datasize", (largeTextLenMax - largeTextLenMin))

	//txt := GetNextText(largeTextLenMin, largeTextLenMax)
	//fmt.Println("txt", txt)
	//fmt.Println("len(txt)", len(txt))

	return GetNextText(largeTextLenMin, largeTextLenMax)
}

func GetNextLargeTextBytes() []byte {
	return []byte(GetNextLargeText())
}
