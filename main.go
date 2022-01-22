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

	"github.com/spf13/viper"
)

type ZipCode struct {
	Zipcode      string `json:"zipcode"`
	Area         string `json:"area"`
	ProvinceCity string `json:"provinceCity"`
}

func main() {
	http.HandleFunc("/zipcode", handleZipcode)

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("ERROR: Error reading config file. Err:", err)
	}
	viper.SetDefault("APP.PORT", "8081")
	PORT, ok := viper.Get("APP.PORT").(string)

	if !ok {
		log.Fatalln("ERROR: Invalid type assertion")
	}

	fmt.Println("INFO: Starting web server at port", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Fatalln("ERROR: Unable to listen and serve, Err:", err)
	}

}

func handleZipcode(rw http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case http.MethodGet:
		err = getZipCode(rw, r)
	}

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func getZipCode(rw http.ResponseWriter, r *http.Request) (err error) {
	start := time.Now()

	if r.URL.Path != "/zipcode" {
		http.Error(rw, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method not supported.", http.StatusNotFound)
		return
	}

	key := r.FormValue("key")

	//open file
	f, err := os.Open("data/zipcode.csv")
	if err != nil {
		log.Fatal(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	//close when program ends
	defer f.Close()

	//read values
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	resultZipcode := []ZipCode{}

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

	result, err := json.Marshal(resultZipcode)
	if err != nil {
		log.Fatal(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	elapsed := time.Since(start)

	logStr := "Key:" + key + ", Result:" + string(result) + ". Completed after " + fmt.Sprint(elapsed.Milliseconds()) + "ms"
	fmt.Println("INFO:", logStr)

	rw.Write(result)
	rw.WriteHeader(http.StatusOK)
	return
}
