package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"image/color"
	png "image/png"
    //"reflect"
    "log"
)

func main() {
	targetDir := "Japanese Handwritten Digits"
    catalogs, err := ioutil.ReadDir(targetDir)
    if err != nil {
        log.Fatal(err)
    }
	i := 0
    for _, catalog := range catalogs {
		if !catalog.IsDir() {
			continue
		}
		
		dirName := (targetDir + "\\" + catalog.Name())

		
		files, _ := ioutil.ReadDir(dirName)
		for _, fileInfo := range files {
			if fileInfo.IsDir() {
				continue
			}

			file, _ := os.Open(dirName + "\\" + fileInfo.Name())
			pngImage, _ := png.Decode(file)
			file.Close()
			
			levels := []string{" ", "█", "█", "█", "█"}
			for y := pngImage.Bounds().Min.Y; y < pngImage.Bounds().Max.Y; y++ {
				for x := pngImage.Bounds().Min.X; x < pngImage.Bounds().Max.X; x++ {
					c := color.GrayModel.Convert(pngImage.At(x, y)).(color.Gray)
					level := c.Y / 51 // 51 * 5 = 255
					if level == 5 {
						level--
					}
					fmt.Print(levels[level])
				}
				fmt.Print("\n")
			}

			/*body, _ := ioutil.ReadFile(dirName + "\\" + fileInfo.Name())
			fmt.Println(body)*/


			break	
		}
		i++
		if i > 9 {
			return
		}
    }
}