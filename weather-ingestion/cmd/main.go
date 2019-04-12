package main

import (
	"os"
	"encoding/csv"
	"io/ioutil"
	"fmt"
	"log"
)

// username = doadmin
// password = yfo38ax5fttdieee
// host = odyssey-pg-test-do-user-2339835-0.db.ondigitalocean.com
// port = 25060
// database = defaultdb
// sslmode = require

// postgres://doadmin:yfo38ax5fttdieee@odyssey-pg-test-do-user-2339835-0.db.ondigitalocean.com:25060/defaultdb?sslmode=require

func main() {

	files, err := ioutil.ReadDir("./gsom-latest")
	if err != nil {
		log.Fatal(err)
	}

	res, err := readCsvFromFile(files[len(files)- 100])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res[0])
	fmt.Println(res[len(res) - 1])
	// for _, f := range files {
	//
	// }



}

func readCsvFromFile(file os.FileInfo) ([][]string, error){
	f, err := os.Open("./gsom-latest/" + file.Name())
	if err != nil {
		return nil, err
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}