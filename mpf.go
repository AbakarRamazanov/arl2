package main

import (
	"os"
	"fmt"
	"strconv"
	"image/png"
	"image/color"
	"encoding/csv"
	"bufio"
)

var numbers = []string{
	"zero", "one", "two", "three", "four",
	"five", "six", "seven", "eight", "nine"}

func sHemming(str1, str2 string) float64 {
	p := 0
	for i, _ := range str1 {
		if (str1[i] != str2[i]) {
			p++
		}
	}
	return float64(p)
}

func sHemmingForClassFromCSV(fileName string, str2 string) float64 {
	var e float64
	e = 0

	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	line, _ := reader.Read()
	for _, str1 := range line{
		e = e + (1/(sHemming(str1, str2) + 1))
	}
	csvFile.Close()

	return e
}

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
	var sumArray [10]float64
	var max float64
	var number int
	str2 := PngTo01InString(os.Args[1])
	for i := 0; i < 10; i++ {
		sumArray[i] = sHemmingForClassFromCSV(("CSVs\\" + strconv.Itoa(i) + ".csv"), str2)
		// return
	}
	for i := 0; i < 10; i++ {
		if (max < sumArray[i]) {
			max = sumArray[i]
			number = i
		}
		fmt.Println(strconv.Itoa(i) + " = "+ fmt.Sprintf("%f",sumArray[i]))
	}
	
	fmt.Println("It is " + numbers[number] + "!")
}
