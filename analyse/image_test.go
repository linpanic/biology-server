package analyse

import (
	"bytes"
	"fmt"
	"github.com/linpanic/biology-server/cst"
	"github.com/linpanic/biology-server/logs"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"testing"
)

func TestImage(t *testing.T) {
	logs.LogInit()

	fp := "./cal/image1.png"
	// 打开文件
	file, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	var img image.Image
	img, err = jpeg.Decode(file)
	if err != nil {
		_ = file.Close()
		file, err = os.Open(fp)
		img, err = png.Decode(file)
		if err != nil {
			panic(err)
		} else {
			_ = file.Close()
		}
	} else {
		defer file.Close()
	}

	var red color.RGBA
	red.R = 89

	newImg := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx(), img.Bounds().Dy()))

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			dis := math.Sqrt(math.Pow(float64(r>>8)-float64(red.R), 2) + math.Pow(float64(g>>8)-float64(red.G), 2) + math.Pow(float64(b>>8)-float64(red.B), 2))
			f := 1 - dis/math.Sqrt(3*math.Pow(255, 2))
			if f > 0.85 {
				newImg.Set(x, y, img.At(x, y))
			} else {
				newImg.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}

	newFn := "./gen.png"
	create, _ := os.Create(newFn)
	defer create.Close()

	err = png.Encode(create, newImg)
	if err != nil {
		panic(err)
	}

}

func TestImg(t *testing.T) {
	filePath := "../image1.png"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// 解码文件，确定其格式
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}
	file.Seek(0, 0)

	bs := make([]byte, 10)
	_, err = file.Read(bs)
	fmt.Println(string(bs))

	switch format {
	case "jpeg":
		fmt.Println("The file is a JPEG image.")
	case "png":
		fmt.Println("The file is a PNG image.")
	default:
		fmt.Println("The file is neither a PNG nor a JPEG image. It is a", format, "image.")
	}
}

func GetImageType(imgBytes []byte) int {
	switch {
	case bytes.Contains(imgBytes, cst.JPEG_TAG):
		return cst.JPEG_TYPE
	case bytes.Contains(imgBytes, cst.PNG_TAG):
		return cst.PNG_TYPE
	}
	return 0
}

func TestGetStrong(t *testing.T) {
	s1, err := GetStronger(AnalyseImage{
		ImagePath:   "../image3.jpg",
		TargetColor: color.RGBA{R: 65},
	})
	if err != nil {
		panic(err)
	}

	s2, err := GetStronger(AnalyseImage{
		ImagePath:   "../image4.jpg",
		TargetColor: color.RGBA{R: 65},
	})
	if err != nil {
		panic(err)
	}

	m1 := make([]Point, len(s1)*len(s1[0]))
	for i, v := range s1 {
		for i2, v2 := range v {
			m1[i*len(v)+i2] = Point{Strong: v2}
		}
	}
	occurrence := GetCoOccurrence(s1, s2, 30, 40)
	fmt.Println(occurrence)
}

type Point struct {
	Strong float64
}
