package setup

import (
	"bytes"
	"github.com/rbo-17/95737-final-project/utils"
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

	img, _, err = image.Decode(f)
	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()
	width = bounds.Dx()
	height = bounds.Dy()
}

func GetNextImage(opts utils.TestOpts) image.Image {

	xStartRange := width - imgWidthMax
	xStart := GetRandRange(0, xStartRange)

	yStartRange := height - imgWidthMax
	yStart := GetRandRange(0, yStartRange)

	newWidth := GetRandRange(imgWidthMin, imgWidthMax)

	cropSize := image.Rect(0, 0, newWidth, newWidth)
	cropSize = cropSize.Add(image.Point{xStart, yStart})
	croppedImage := img.(SubImager).SubImage(cropSize)

	return croppedImage
}

func GetNextImageBytes(opts utils.TestOpts) ([]byte, error) {

	var res []byte

	// Create n independent images and append the bytes. TODO: Consider creating one large image instead
	for i := 0; i < opts.DenormalizationFactor; i++ {
		nextImage := GetNextImage(opts)

		buf := new(bytes.Buffer)
		err := png.Encode(buf, nextImage)
		if err != nil {
			return nil, err
		}

		res = append(res, buf.Bytes()...)
	}

	return res, nil
}
