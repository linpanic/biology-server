package analyse

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

type AnalyseImage struct {
	ImagePath   string     //图片路径
	TargetColor color.RGBA //目标RGB A忽略
	Tolerange   float64    //容差，相似度在1-容差内的话，则符合条件
}

type ImageStrong struct {
	Strong [][]float64
}

func GetStronger(i AnalyseImage) ([][]float64, error) {
	file, err := os.Open(i.ImagePath)
	if err != nil {
		return nil, err
	}

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}
	file.Seek(0, 0)
	var img image.Image
	switch format {
	case "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	default:
		return nil, errors.New("The file is neither a PNG nor a JPEG image. It is a " + format + "image")
	}
	if err != nil {
		return nil, err
	}

	sts := make([][]float64, img.Bounds().Dx())
	for idx := range sts {
		sts[idx] = make([]float64, img.Bounds().Dy())
	}

	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			var strong float64
			r32, g32, b32, _ := img.At(x, y).RGBA()
			switch {
			case i.TargetColor.R != 0 && r32 > 255:
				strong = float64(i.TargetColor.R / uint8(r32>>8))
			case i.TargetColor.G != 0 && g32 > 255:
				strong = float64(i.TargetColor.G / uint8(g32>>8))
			case i.TargetColor.B != 0 && b32 > 255:
				strong = float64(i.TargetColor.B / uint8(b32>>8))
			}
			sts[x][y] = strong
		}
	}
	return sts, nil
}

func GetCoOccurrence(iA, iB [][]float64, tA, tB float64) float64 {
	var numerator, denominator float64

	for x := 0; x < len(iA); x++ {
		for y := 0; y < len(iA[x]); y++ {
			pA := iA[x][y]
			if pA > tA {
				denominator += pA
				pB := iB[x][y]
				if pB > tB {
					numerator += pA
				}
			}
		}
	}

	return numerator / denominator
}
