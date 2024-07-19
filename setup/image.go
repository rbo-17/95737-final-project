package setup

import (
	"bytes"
	"image"
	"image/png"
	"os"
)

const ImagePath = "data/school-of-athens.png"

const imgWidthMin = 815  // 1MB / 0.2205 multiplier
const imgWidthMax = 1129 // 1.99MB / 0.3054 multiplier

var width, height int
var img image.Image

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func init() {
	f, err := os.Open(ImagePath)
	if err != nil {
		panic(err)
	}

	//fmt.Println(ImagePath)

	img, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
	//fmt.Println("width", width)
	//fmt.Println("height", height)
}

func GetNextImage() image.Image {

	xStartRange := width - imgWidthMax
	xStart := GetRandRange(0, xStartRange)

	yStartRange := height - imgWidthMax
	yStart := GetRandRange(0, yStartRange)

	newWidth := GetRandRange(imgWidthMin, imgWidthMax)

	//fmt.Println("img", img)
	cropSize := image.Rect(0, 0, newWidth, newWidth)
	cropSize = cropSize.Add(image.Point{xStart, yStart})
	croppedImage := img.(SubImager).SubImage(cropSize)

	return croppedImage
}

func GetNextImageBytes() []byte {

	nextImage := GetNextImage()

	buf := new(bytes.Buffer)
	err := png.Encode(buf, nextImage)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func PrintImage() {
	//nf, err := os.Create("image.png")
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = png.Encode(nf, croppedImage)
	//if err != nil {
	//	panic(err)
	//}
}
