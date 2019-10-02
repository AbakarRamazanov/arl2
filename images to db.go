package main

import (
	"os"
	//"fmt"
	"io/ioutil"
	"image/color"
	"image/png"
	//"reflect"
    "encoding/csv"
    "log"
)

func PngTo01InString(fileName string) string {
	file, _ := os.Open(fileName)
	defer file.Close()
	pngImage, _ := png.Decode(file)
	pngString := ""
	for y := pngImage.Bounds().Min.Y; y < pngImage.Bounds().Max.Y; y++ {
		for x := pngImage.Bounds().Min.X; x < pngImage.Bounds().Max.X; x++ {	
			c := color.GrayModel.Convert(pngImage.At(x, y)).(color.Gray)
			if c.Y >= 10 {
				pngString += "1"
			} else {
				pngString += "0"
			}
		}
	}
	return pngString
}

func main() {
	targetDir := "Japanese Handwritten Digits"
    catalogs, err := ioutil.ReadDir(targetDir)
    if err != nil {
        log.Fatal(err)
    }
    for _, catalog := range catalogs {
		if !catalog.IsDir() {
			continue
		}		
		dirName := (targetDir + "\\" + catalog.Name())
		files, _ := ioutil.ReadDir(dirName)
		pngStrings := make([]string, 0)
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				continue
			}
			pngStrings =  append( pngStrings, PngTo01InString(dirName + "\\" + fileInfo.Name()))
		}
		fileCsv, _ := os.Create(catalog.Name() + ".csv")
		writer := csv.NewWriter(fileCsv)
		writer.Write(pngStrings);
		writer.Flush()
		fileCsv.Close()
    }
}
