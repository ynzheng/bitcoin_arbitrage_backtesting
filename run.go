package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	/////////////////////////////////////////////////////////
	//           Datasource handeler
	/////////////////////////////////////////////////////////
	// Datasource filename
	var 

	// Read current directory
	runningDir, runningDirErr := filepath.Abs(filepath.Dir(os.Args[0]))
	if runningDirErr != nil {
		log.Fatal(runningDirErr)
	}
	fmt.Println(runningDir)

	// Open file and create GZ reader
	f, err := os.Open(runningDir + "/data/bitfinexUSD.csv.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()

	// Read CSV from GZ file
	cr := csv.NewReader(gr)
	rec, err := cr.Read()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range rec {
		fmt.Print(v, " ")
	}

	// Datetime parser
	/*i, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)*/


	/////////////////////////////////////////////////////////
	//           ARBITRAGE Calculation
	/////////////////////////////////////////////////////////

}
