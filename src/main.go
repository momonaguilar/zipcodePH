package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type ZipCode struct {
	Zipcode      string
	Area         string
	ProvinceCity string
}

func main() {
	http.HandleFunc("/getZip", handleGetZip)
	fmt.Println("Starting web server at port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

func handleGetZip(rw http.ResponseWriter, r *http.Request) {
	//time start
	start := time.Now()

	if r.URL.Path != "/getZip" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method not supported.", http.StatusNotFound)
		return
	}

	search := r.FormValue("search")

	//open file
	f, err := os.Open("zipcode.csv")
	if err != nil {
		log.Fatal(err)
	}

	//close when program ends
	defer f.Close()

	//read values
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var resultZipcode []ZipCode

	for _, line := range data {
		if line[0] == search || line[1] == search || line[2] == search {
			z := ZipCode{
				Zipcode:      line[0],
				Area:         line[1],
				ProvinceCity: line[2],
			}
			resultZipcode = append(resultZipcode, z)
		}
	}

	result, err := json.Marshal(resultZipcode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result))

	elapsed := time.Since(start)
	fmt.Println("Done after: ", elapsed.Milliseconds(), "ms")

}
