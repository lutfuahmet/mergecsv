package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

// mergecsv sample1.csv sample2.csv out.csv

var sources []string
var target string

var checkMap = map[string]bool{}

func main() {
	log.Println(os.Args)
	initFiles()
	readFiles()
}

func readFiles()  {
	targetCsv,err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	defer targetCsv.Close()
	csvwriter := csv.NewWriter(targetCsv)
	for _, source := range sources {
		lines,err := readFile(source)
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range lines {
			firstCell := line[0]
			if !checkMap[firstCell]{
				csvwriter.Write(line)
				checkMap[firstCell] = true
			}
		}
	}
	csvwriter.Flush()
}

func initFiles()  {
	var files []string
	for _, arg := range os.Args {
		if strings.HasSuffix(arg,".csv"){
			files = append(files,arg)
		}
	}
	if len(files) == 0 {
		return
	}
	sources = files[:len(files) - 1]
	target = files[len(files) - 1] // last param
	log.Println(sources)
	log.Println(target)
}

func readFile(fileName string) ([][]string,error)  {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil,err
	}
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil,err
	}
	for _,line := range csvLines {
		for _, cell := range line {
			log.Println(cell)
		}
	}
	return csvLines,nil
}