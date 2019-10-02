package main

import (
	"os"
    "encoding/csv"
	"strconv"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	for i := 0; i < 10; i++ {		
		csvFile, _ := os.Open(strconv.Itoa(i) + ".csv")
		reader := csv.NewReader(bufio.NewReader(csvFile))
		line, _ := reader.Read()
		stringArray:=strings.Split(line[0], " ")
		fmt.Println()
		for _, img := range stringArray{
			for o1 := 0; o1 < 28; o1++ {
				for o2 := 0; o2 < 28; o2++ {
					fmt.Print(img[o1*28+o2]-48)
				}
				fmt.Println()
			}
			return
		}
		csvFile.Close()
		//fmt.Println(line)
		return
	}
}