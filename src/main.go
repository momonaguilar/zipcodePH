package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type ZipCode struct {
	Zipcode      string
	Area         string
	ProvinceCity string
}

func main() {
	http.HandleFunc("/zipcode", handleGetZip)
	fmt.Println("Starting web server at port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func handleGetZip(rw http.ResponseWriter, r *http.Request) {
	//time start
	start := time.Now()

	fmt.Fprintln(rw, "INFO: Starting")

	if r.URL.Path != "/zipcode" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	fmt.Fprintln(rw, "INFO: Starting")

	if r.Method != "GET" {
		http.Error(rw, "Method not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintln(rw, "INFO: Starting")

	key := r.FormValue("key")

	//open file
	f, err := os.Open("zipcode.csv")
	if err != nil {
		log.Fatal(err)
	}

	//close when program ends
	defer f.Close()

	fmt.Fprintln(rw, "INFO: Starting")

	//read values
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(rw, key)
	fmt.Fprintln(rw, "INFO: Starting...")
	var resultZipcode []ZipCode

	for _, line := range data {
		if line[0] == key || strings.EqualFold(line[1], key) {
			z := ZipCode{
				Zipcode:      line[0],
				Area:         line[1],
				ProvinceCity: line[2],
			}
			resultZipcode = append(resultZipcode, z)
		}
	}

	fmt.Fprintln(rw, "INFO: Starting")
	result, err := json.Marshal(resultZipcode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(rw, string(result))

	elapsed := time.Since(start)
	fmt.Println("Done after: ", elapsed.Milliseconds(), "ms")

}
