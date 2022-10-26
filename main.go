package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

var dataTXT string
var dataHTML string
var comma string
var max int = 0
var lenAll [][]string
var PathToCsvFile string = "1.csv"

func main() {
	settings()
	max_len()
	lenAllFile()
	read_csv()
	writer()
}

func settings() {
	fmt.Println("Please, enter a text separator: ")
	fmt.Fscan(os.Stdin, &comma)
}

func max_len() {
	file1, err := os.Open(PathToCsvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file1.Close()
	reader := csv.NewReader(file1)
	reader.Comma = rune(comma[0])
	
	for {
		lenR, e := reader.Read()
		if e != nil {
			fmt.Println(e)
		}
		for i := range lenR {
			
			if max < len(lenR[i]) {
				max = len(lenR[i])
			}
			
		}
		if e != nil {
			fmt.Println(e)
			break
		}
	}
	fmt.Print(max)
}

func lenAllFile() {
	file1, err := os.Open(PathToCsvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file1.Close()
	reader := csv.NewReader(file1)
	reader.Comma = rune(comma[0])

	lA, e := reader.ReadAll()

	if e != nil {
		fmt.Println(e)
	}
	lenAll = lA
}

func read_csv() {
	fmt.Println("Reader:")
	file, err := os.Open(PathToCsvFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = rune(comma[0])

	
	n := 0
	n1 := 0
	

	for {
		record, e := reader.Read()
		fmt.Println(record)
		if n == 0 {
			dataHTML += "<table>\n    <tr>"
			dataTXT += strings.Repeat("+"+strings.Repeat("-", max), len(record)) + "+\n"
		}
		n = 1
		n1 += 1
		for j := range record {
			d := "|" + record[j]
			if len(d) < max+1 {
				d += strings.Repeat(" ", (max+1)-len(d))
				fmt.Println(max)
			}
			dataTXT += d
			dataHTML += "\n        <td>" + record[j] + "</td>"
		}
		
		if n1 < len(lenAll) {
			dataHTML += "\n    </tr>\n    <tr>"
		}
		if n1 <= len(lenAll) {
			dataTXT += "|\n" + strings.Repeat("+"+strings.Repeat("-", max), len(record)) + "+\n"
		}
		if e != nil {
			fmt.Println(e)
			break
		}
	}
	dataHTML += "\n    </tr>\n</table>"
}

func writer() {
	fmt.Println("Writer:")
	file, err := os.Create("spreadsheet.txt")
	text := []byte(dataTXT)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(text)

	file1, err := os.Create("spreadsheet.html")
	text1 := []byte(dataHTML)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file1.Close()
	file1.Write(text1)
}
