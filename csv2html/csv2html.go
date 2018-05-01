package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	// "flag"
	// "path"
	// "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO: Add flags rather than
// position based arguments
// TODO: Add Style
// 			 Add markdown format

func main() {
	var fileIn, fileOut string
	args := os.Args[1:]
	numArgs := len(args)
	fmt.Println(numArgs)
	switch numArgs {
	case 1:
		fileIn = args[0]
	case 2:
		fileIn = args[0]
		fileOut = args[1]
	}

	working := "<head><link rel='stylesheet' type='text/css' href='./css/dist/milligram.min.css'></head>\n"
	working += "<body>\n"
	// path := path.Base(fileIn)

	// file_chunks := strings.Split(path, ".")
	// fileOut := file_chunks[0] + ".html"
	// fmt.Println(file_chunks)

	f, err := os.Open(fileIn)
	check(err)
	r := csv.NewReader(bufio.NewReader(f))
	rowNum := 1
	working += "<table>\n"
	for {
		record, err := r.Read()
		if rowNum == 1 {
			working += "<th>"
			for _, v := range record {
				working += "<td>" + v + "</td>\n"
			}
			working += "</th>"
		} else {
			working += "<tr>"
			for _, v := range record {
				working += "<td>" + v + "</td>\n"
			}
			working += "</tr>"
		}
		rowNum++

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(record)
	}
	working += "</table>\n"
	working += "</body>\n"

	if numArgs > 1 {
		out, err := os.Create(fileOut)
		check(err)
		w := bufio.NewWriter(out)
		_, err = w.WriteString(working)
		check(err)
		w.Flush()
	} else {
		fmt.Println(working)
	}
}
